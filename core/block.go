// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/nebulasio/go-nebulas/common/dag"
	"github.com/nebulasio/go-nebulas/common/dag/pb"
	"github.com/nebulasio/go-nebulas/common/mvccdb"
	"github.com/nebulasio/go-nebulas/consensus/pb"
	"github.com/nebulasio/go-nebulas/core/pb"
	"github.com/nebulasio/go-nebulas/core/state"
	"github.com/nebulasio/go-nebulas/crypto"
	"github.com/nebulasio/go-nebulas/crypto/keystore"
	"github.com/nebulasio/go-nebulas/storage"
	"github.com/nebulasio/go-nebulas/util"
	"github.com/nebulasio/go-nebulas/util/byteutils"
	"github.com/nebulasio/go-nebulas/util/logging"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

var (
	// BlockHashLength define a const of the length of Hash of Block in byte.
	BlockHashLength = 32

	// BlockReward given to coinbase
	// rule: 3% per year, 3,000,000. 1 block per 5 seconds
	// value: 10^8 * 3% / (365*24*3600/5) * 10^18 ≈ 16 * 3% * 10*18 = 48 * 10^16
	BlockReward, _ = util.NewUint128FromString("480000000000000000") // TODO 4x in 20s
)

// BlockHeader of a block
type BlockHeader struct {
	hash       byteutils.Hash
	parentHash byteutils.Hash

	// world state
	stateRoot     byteutils.Hash
	txsRoot       byteutils.Hash
	eventsRoot    byteutils.Hash
	consensusRoot *consensuspb.ConsensusRoot

	coinbase  *Address
	timestamp int64
	chainID   uint32

	// sign
	alg  keystore.Algorithm
	sign byteutils.Hash
}

// ToProto converts domain BlockHeader to proto BlockHeader
func (b *BlockHeader) ToProto() (proto.Message, error) {
	return &corepb.BlockHeader{
		Hash:          b.hash,
		ParentHash:    b.parentHash,
		StateRoot:     b.stateRoot,
		TxsRoot:       b.txsRoot,
		EventsRoot:    b.eventsRoot,
		ConsensusRoot: b.consensusRoot,
		Coinbase:      b.coinbase.address,
		Timestamp:     b.timestamp,
		ChainId:       b.chainID,
		Alg:           uint32(b.alg),
		Sign:          b.sign,
	}, nil
}

// FromProto converts proto BlockHeader to domain BlockHeader
func (b *BlockHeader) FromProto(msg proto.Message) error {
	if msg, ok := msg.(*corepb.BlockHeader); ok {
		b.hash = msg.Hash
		b.parentHash = msg.ParentHash
		b.stateRoot = msg.StateRoot
		b.txsRoot = msg.TxsRoot
		b.eventsRoot = msg.EventsRoot
		if msg.ConsensusRoot == nil {
			return ErrInvalidProtoToBlockHeader
		}
		b.consensusRoot = msg.ConsensusRoot
		coinbase, err := AddressParseFromBytes(msg.Coinbase)
		if err != nil {
			return ErrInvalidProtoToBlockHeader
		}
		b.coinbase = coinbase
		b.timestamp = msg.Timestamp
		b.chainID = msg.ChainId
		b.alg = keystore.Algorithm(msg.Alg) // TODO: check category
		b.sign = msg.Sign
		return nil
	}
	return ErrInvalidProtoToBlockHeader
}

// Block structure
type Block struct {
	header       *BlockHeader
	transactions Transactions
	dependency   *dag.Dag

	sealed      bool
	height      uint64
	parentBlock *Block

	worldState state.WorldState

	txPool       *TransactionPool
	eventEmitter *EventEmitter
	nvm          NVM
	storage      storage.Storage
}

// ToProto converts domain Block into proto Block
func (block *Block) ToProto() (proto.Message, error) {
	header, err := block.header.ToProto()
	if err != nil {
		return nil, err
	}
	if header, ok := header.(*corepb.BlockHeader); ok {
		txs := make([]*corepb.Transaction, len(block.transactions))
		for idx, v := range block.transactions {
			tx, err := v.ToProto()
			if err != nil {
				return nil, err
			}
			if tx, ok := tx.(*corepb.Transaction); ok {
				txs[idx] = tx
			} else {
				return nil, ErrInvalidProtoToTransaction
			}
		}
		dependency, err := block.dependency.ToProto()
		if err != nil {
			return nil, err
		}
		if dependency := dependency.(*dagpb.Dag); ok {
			return &corepb.Block{
				Header:       header,
				Transactions: txs,
				Dependency:   dependency,
				Height:       block.height,
			}, nil
		}
		return nil, ErrInvalidProtoToDag
	}
	return nil, ErrInvalidProtoToBlock
}

// FromProto converts proto Block to domain Block
func (block *Block) FromProto(msg proto.Message) error {
	if msg, ok := msg.(*corepb.Block); ok {
		block.header = new(BlockHeader)
		if err := block.header.FromProto(msg.Header); err != nil {
			return err
		}

		block.transactions = make(Transactions, len(msg.Transactions))
		for idx, v := range msg.Transactions {
			tx := new(Transaction)
			if err := tx.FromProto(v); err != nil {
				return err
			}
			block.transactions[idx] = tx
		}
		block.dependency = dag.NewDag()
		if err := block.dependency.FromProto(msg.Dependency); err != nil { // TODO: check nil in all FromProto, add unit tests
			return err
		}
		block.height = msg.Height
		return nil
	}
	return ErrInvalidProtoToBlock
}

// NewBlock return new block.
func NewBlock(chainID uint32, coinbase *Address, parent *Block) (*Block, error) { // ToCheck: check args. // ToCheck: check full-functional block.
	worldState, err := parent.worldState.Clone()
	if err != nil {
		return nil, err
	}

	block := &Block{
		header: &BlockHeader{
			chainID:       chainID,
			parentHash:    parent.Hash(),
			coinbase:      coinbase,
			timestamp:     time.Now().Unix(),
			consensusRoot: &consensuspb.ConsensusRoot{},
		},
		transactions: make(Transactions, 0),
		dependency:   dag.NewDag(),
		parentBlock:  parent,

		worldState: worldState,
		height:     parent.height + 1,
		sealed:     false,

		txPool:       parent.txPool,
		eventEmitter: parent.eventEmitter,
		nvm:          parent.nvm,
		storage:      parent.storage,
	}

	if err := block.Begin(); err != nil {
		return nil, err
	}
	if err := block.rewardCoinbaseForMint(); err != nil {
		return nil, err
	}

	return block, nil
}

// Sign sign transaction,sign algorithm is
func (block *Block) Sign(signature keystore.Signature) error { // TODO: check nil
	sign, err := signature.Sign(block.header.hash)
	if err != nil {
		return err
	}
	block.header.alg = keystore.Algorithm(signature.Algorithm())
	block.header.sign = sign
	return nil
}

// ChainID returns block's chainID
func (block *Block) ChainID() uint32 {
	return block.header.chainID
}

// Coinbase return block's coinbase
func (block *Block) Coinbase() *Address {
	return block.header.coinbase
}

// Alg return block's alg
func (block *Block) Alg() keystore.Algorithm {
	return block.header.alg
}

// Signature return block's signature
func (block *Block) Signature() byteutils.Hash {
	return block.header.sign
}

// Timestamp return timestamp
func (block *Block) Timestamp() int64 {
	return block.header.timestamp
}

// SetTimestamp set timestamp
func (block *Block) SetTimestamp(timestamp int64) {
	if block.sealed {
		logging.VLog().WithFields(logrus.Fields{
			"block": block,
		}).Fatal("Sealed block can't be changed.")
	}
	block.header.timestamp = timestamp
}

// Hash return block hash.
func (block *Block) Hash() byteutils.Hash {
	return block.header.hash
}

// StateRoot return state root hash.
func (block *Block) StateRoot() byteutils.Hash {
	return block.header.stateRoot
}

// TxsRoot return txs root hash.
func (block *Block) TxsRoot() byteutils.Hash {
	return block.header.txsRoot
}

// Storage return storage.
func (block *Block) Storage() storage.Storage {
	return block.storage
}

// WorldState return the world state of the block
func (block *Block) WorldState() state.WorldState {
	return block.worldState
}

// EventsRoot return events root hash.
func (block *Block) EventsRoot() byteutils.Hash {
	return block.header.eventsRoot
}

// ConsensusRoot return consensus root
func (block *Block) ConsensusRoot() *consensuspb.ConsensusRoot {
	return block.header.consensusRoot
}

// ParentHash return parent hash.
func (block *Block) ParentHash() byteutils.Hash {
	return block.header.parentHash
}

// Height return height
func (block *Block) Height() uint64 {
	return block.height
}

// Transactions returns block transactions
func (block *Block) Transactions() Transactions {
	return block.transactions
}

// LinkParentBlock link parent block, return true if hash is the same; false otherwise.
func (block *Block) LinkParentBlock(chain *BlockChain, parentBlock *Block) error {
	if !block.ParentHash().Equals(parentBlock.Hash()) {
		return ErrLinkToWrongParentBlock
	}

	var err error
	if block.worldState, err = parentBlock.WorldState().Clone(); err != nil {
		return ErrCloneAccountState
	}

	elapsedSecond := block.Timestamp() - parentBlock.Timestamp()
	consensusState, err := parentBlock.worldState.NextConsensusState(elapsedSecond)
	if err != nil {
		return err
	}
	block.WorldState().SetConsensusState(consensusState)

	block.height = parentBlock.height + 1
	block.parentBlock = parentBlock
	block.txPool = parentBlock.txPool
	block.storage = parentBlock.storage
	block.eventEmitter = parentBlock.eventEmitter
	block.nvm = parentBlock.nvm

	return nil
}

// Begin a batch task
func (block *Block) Begin() error {
	return block.WorldState().Begin()
}

// Commit a batch task
func (block *Block) Commit() {
	if err := block.WorldState().Commit(); err != nil {
		logging.CLog().Fatal(err)
	}
}

// RollBack a batch task
func (block *Block) RollBack() {
	if err := block.WorldState().RollBack(); err != nil {
		logging.CLog().Fatal(err)
	}
}

// ReturnTransactions and giveback them to tx pool
// TODO(roy): optimize storage.
// if a block is reverted, we should erase all changes
// made by this block on storage. use refcount.
func (block *Block) ReturnTransactions() {
	for _, tx := range block.transactions {
		block.txPool.Push(tx)
	}
}

// CollectTransactions and add them to block.
func (block *Block) CollectTransactions(deadlineInMs int64) {
	metricsBlockPackTxTime.Update(0)
	if block.sealed {
		logging.VLog().WithFields(logrus.Fields{
			"block": block,
		}).Fatal("Sealed block can't be changed.")
	}

	secondInMs := int64(1000)
	elapseInMs := deadlineInMs - time.Now().Unix()*secondInMs
	logging.VLog().WithFields(logrus.Fields{
		"elapse": elapseInMs,
	}).Info("Time to pack txs.")
	metricsBlockPackTxTime.Update(elapseInMs)
	if elapseInMs <= 0 {
		return
	}
	deadlineTimer := time.NewTimer(time.Duration(elapseInMs) * time.Millisecond)

	pool := block.txPool

	packed := int64(0)
	unpacked := int64(0)

	dag := dag.NewDag()
	transactions := []*Transaction{}
	fromBlacklist := new(sync.Map)
	toBlacklist := new(sync.Map)

	parallelCh := make(chan bool, 32)
	mergeCh := make(chan bool, 1) // TODO: add comments in all usage of mergeCh
	over := false

	try := 0
	fetch := 0
	failed := 0
	conflict := 0
	expired := 0
	bucket := len(block.txPool.all)
	packing := int64(0)
	prepare := int64(0)
	execute := int64(0)
	update := int64(0)
	parallel := 0
	beginAt := time.Now().UnixNano()

	go func() {
		for {
			mergeCh <- true
			if over {
				<-mergeCh
				return
			}
			try++
			tx := pool.PopWithBlacklist(fromBlacklist, toBlacklist)
			if tx == nil {
				<-mergeCh
				// time.Sleep(time.Nanosecond * 1000)
				continue
			}
			fetch++
			fromBlacklist.Store(tx.from.address.Hex(), true)
			fromBlacklist.Store(tx.to.address.Hex(), true)
			toBlacklist.Store(tx.from.address.Hex(), true)
			toBlacklist.Store(tx.to.address.Hex(), true)
			<-mergeCh

			parallelCh <- true
			go func() {
				parallel++
				startAt := time.Now().UnixNano()
				defer func() {
					endAt := time.Now().UnixNano()
					packing += endAt - startAt
					<-parallelCh
				}()

				mergeCh <- true
				if over {
					expired++
					<-mergeCh
					if err := pool.Push(tx); err != nil {
						logging.VLog().WithFields(logrus.Fields{
							"block": block,
							"tx":    tx,
							"err":   err,
						}).Debug("Failed to giveback the tx.")
					}
					return
				}

				prepareAt := time.Now().UnixNano()
				txWorldState, err := block.WorldState().Prepare(tx.Hash().String())
				preparedAt := time.Now().UnixNano()
				prepare += preparedAt - prepareAt
				if err != nil {
					logging.VLog().WithFields(logrus.Fields{
						"block": block,
						"tx":    tx,
						"err":   err,
					}).Debug("Failed to prepare tx.") // TODO: push back tx
					fromBlacklist.Delete(tx.from.address.Hex())
					fromBlacklist.Delete(tx.to.address.Hex())
					toBlacklist.Delete(tx.from.address.Hex())
					toBlacklist.Delete(tx.to.address.Hex())
					<-mergeCh
					return
				}
				<-mergeCh

				executeAt := time.Now().UnixNano()
				giveback, err := block.ExecuteTransaction(tx, txWorldState) // TODO: move giveback logic into Execution
				executedAt := time.Now().UnixNano()
				execute += executedAt - executeAt
				if err != nil {
					logging.CLog().WithFields(logrus.Fields{
						"tx":       tx,
						"err":      err,
						"giveback": giveback,
					}).Debug("invalid tx.")
					unpacked++ // TODO close txWorldState

					if giveback || err == mvccdb.ErrPreparedDBIsClosed {
						if err := pool.Push(tx); err != nil {
							logging.VLog().WithFields(logrus.Fields{
								"block": block,
								"tx":    tx,
								"err":   err,
							}).Debug("Failed to giveback the tx.")
						}
						failed++                                  // TODO move out of if-else
						fromBlacklist.Delete(tx.to.address.Hex()) // TODO: add comment, why not remove tx.from
						toBlacklist.Delete(tx.to.address.Hex())
					} else {
						fromBlacklist.Delete(tx.from.address.Hex())
						fromBlacklist.Delete(tx.to.address.Hex())
						toBlacklist.Delete(tx.from.address.Hex())
						toBlacklist.Delete(tx.to.address.Hex())
					} // TODO: return directly
				} else {
					mergeCh <- true
					if over {
						expired++
						<-mergeCh
						if err := pool.Push(tx); err != nil {
							logging.VLog().WithFields(logrus.Fields{
								"block": block,
								"tx":    tx,
								"err":   err,
							}).Debug("Failed to giveback the tx.")
						}
						return
					}
					updateAt := time.Now().UnixNano()
					dependency, err := txWorldState.CheckAndUpdate()
					updatedAt := time.Now().UnixNano()
					update += updatedAt - updateAt
					if err != nil {
						logging.VLog().WithFields(logrus.Fields{
							"tx":         tx,
							"err":        err,
							"giveback":   giveback,
							"dependency": dependency,
						}).Debug("CheckAndUpdate invalid tx.") // TODO: release mergeCh
						unpacked++

						if err := txWorldState.Close(); err != nil {
							logging.VLog().WithFields(logrus.Fields{
								"block": block,
								"tx":    tx,
								"err":   err,
							}).Debug("Failed to close tx.")
						}
						if err := pool.Push(tx); err != nil {
							logging.VLog().WithFields(logrus.Fields{
								"block": block,
								"tx":    tx,
								"err":   err,
							}).Debug("Failed to giveback the tx.")
						}
						conflict++
						fromBlacklist.Delete(tx.from.address.Hex())
						fromBlacklist.Delete(tx.to.address.Hex())
						toBlacklist.Delete(tx.from.address.Hex())
						toBlacklist.Delete(tx.to.address.Hex()) // TODO: return directly

					} else {
						logging.CLog().WithFields(logrus.Fields{
							"tx": tx,
						}).Debug("packed tx.")
						packed++

						transactions = append(transactions, tx)
						txid := tx.Hash().String()
						dag.AddNode(txid)
						for _, node := range dependency {
							dag.AddEdge(node, txid)
						}
						fromBlacklist.Delete(tx.from.address.Hex())
						fromBlacklist.Delete(tx.to.address.Hex())
						toBlacklist.Delete(tx.from.address.Hex())
						toBlacklist.Delete(tx.to.address.Hex()) // TODO: release mergeCh, return directly
					}
					<-mergeCh
				}
			}()

			if over {
				return
			}
		}
	}()

	<-deadlineTimer.C
	mergeCh <- true
	over = true
	block.transactions = transactions
	block.dependency = dag // TODO: release mergeCh

	overAt := time.Now().UnixNano()
	size := int64(len(block.transactions))
	if size == 0 {
		size = 1
	}
	averPacking := packing / size
	averPrepare := prepare / size
	averExecute := execute / size
	averUpdate := update / size

	logging.CLog().WithFields(logrus.Fields{
		"try":          try,
		"failed":       failed,
		"expired":      expired,
		"conflict":     conflict,
		"fetch":        fetch,
		"bucket":       bucket,
		"averPacking":  averPacking, // TODO: aver -> avg
		"averPrepare":  averPrepare,
		"averExecute":  averExecute,
		"averUpdate":   averUpdate,
		"parallel":     parallel,
		"packing":      packing,
		"execute":      execute,
		"prepare":      prepare,
		"update":       update,
		"diff-all":     overAt - beginAt,
		"core-packing": execute + prepare + update,
		"packed":       len(block.transactions),
	}).Info("CollectTransactions")

	<-mergeCh
}

// Sealed return true if block seals. Otherwise return false.
func (block *Block) Sealed() bool {
	return block.sealed
}

// Seal seal block, calculate stateRoot and block hash.
func (block *Block) Seal() error {
	if block.sealed {
		logging.VLog().WithFields(logrus.Fields{
			"block": block,
		}).Fatal("cannot seal a block twice.")
	}

	if err := block.rewardCoinbaseForGas(); err != nil {
		return err // TODO: giveback txs
	}
	if err := block.WorldState().Flush(); err != nil {
		return err // TODO: giveback txs
	}
	block.header.stateRoot = block.WorldState().AccountsRoot()
	block.header.txsRoot = block.WorldState().TxsRoot()
	block.header.eventsRoot = block.WorldState().EventsRoot()
	block.header.consensusRoot = block.WorldState().ConsensusRoot()

	var err error
	block.header.hash, err = HashBlock(block)
	if err != nil {
		return err
	}
	block.sealed = true

	logging.VLog().WithFields(logrus.Fields{
		"block": block,
	}).Info("Sealed Block.")

	block.RollBack() // TODO: defer, rollback

	metricsTxPackedCount.Update(0)
	metricsTxUnpackedCount.Update(0)
	metricsTxGivebackCount.Update(0)

	return nil
}

func (block *Block) String() string {
	return fmt.Sprintf(`{"height": %d, "hash": "%s", "parent_hash": "%s", "acc_root": "%s", "timestamp": %d, "tx": %d, "miner": "%s"}`,
		block.height,
		block.header.hash,
		block.header.parentHash,
		block.header.stateRoot,
		block.header.timestamp,
		len(block.transactions),
		byteutils.Hex(block.header.consensusRoot.Proposer),
	)
}

// VerifyExecution execute the block and verify the execution result.
func (block *Block) VerifyExecution() error {
	startAt := time.Now().Unix()

	if err := block.Begin(); err != nil {
		return err
	}

	beganAt := time.Now().Unix()

	if err := block.execute(); err != nil {
		block.RollBack()
		return err
	}

	executedAt := time.Now().Unix()

	if err := block.verifyState(); err != nil {
		block.RollBack()
		return err
	}

	commitAt := time.Now().Unix()

	block.Commit()

	endAt := time.Now().Unix()

	logging.CLog().WithFields(logrus.Fields{
		"start":        startAt,
		"end":          endAt,
		"commit":       commitAt,
		"diff-all":     endAt - startAt,
		"diff-commit":  endAt - commitAt,
		"diff-begin":   beganAt - startAt,
		"diff-execute": executedAt - startAt,
		"diff-verify":  commitAt - executedAt,
		"block":        block,
		"txs":          len(block.Transactions()),
	}).Info("Verify txs.")

	return nil
}

// VerifyIntegrity verify block's hash, txs' integrity and consensus acceptable.
func (block *Block) VerifyIntegrity(chainID uint32, consensus Consensus) error {

	if consensus == nil {
		metricsInvalidBlock.Inc(1)
		return ErrNilArgument
	}

	// check ChainID.
	if block.header.chainID != chainID {
		logging.VLog().WithFields(logrus.Fields{
			"expect": chainID,
			"actual": block.header.chainID,
		}).Debug("Failed to check chainid.")
		metricsInvalidBlock.Inc(1)
		return ErrInvalidChainID
	}

	// verify block hash.
	wantedHash, err := HashBlock(block)
	if err != nil {
		return err
	}
	if !wantedHash.Equals(block.Hash()) {
		logging.VLog().WithFields(logrus.Fields{
			"expect": wantedHash,
			"actual": block.Hash(),
			"err":    err,
		}).Debug("Failed to check block's hash.")
		metricsInvalidBlock.Inc(1)
		return ErrInvalidBlockHash
	}

	// verify transactions integrity.
	for _, tx := range block.transactions {
		if err := tx.VerifyIntegrity(block.header.chainID); err != nil {
			logging.VLog().WithFields(logrus.Fields{
				"tx":  tx,
				"err": err,
			}).Debug("Failed to verify tx's integrity.")
			metricsInvalidBlock.Inc(1)
			return err
		}
	}

	// verify the block is acceptable by consensus.
	if err := consensus.VerifyBlock(block); err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"block": block,
			"err":   err,
		}).Debug("Failed to verify block.")
		metricsInvalidBlock.Inc(1)
		return err
	}

	return nil
}

// verifyState return state verify result.
func (block *Block) verifyState() error {
	// verify state root.
	if !byteutils.Equal(block.WorldState().AccountsRoot(), block.StateRoot()) {
		logging.VLog().WithFields(logrus.Fields{
			"expect": block.StateRoot(),
			"actual": block.WorldState().AccountsRoot(),
		}).Debug("Failed to verify state.")
		return ErrInvalidBlockStateRoot
	}

	// verify transaction root.
	if !byteutils.Equal(block.WorldState().TxsRoot(), block.TxsRoot()) {
		logging.VLog().WithFields(logrus.Fields{
			"expect": block.TxsRoot(),
			"actual": block.WorldState().TxsRoot(),
		}).Debug("Failed to verify txs.")
		return ErrInvalidBlockTxsRoot
	}

	// verify events root.
	if !byteutils.Equal(block.WorldState().EventsRoot(), block.EventsRoot()) {
		logging.VLog().WithFields(logrus.Fields{
			"expect": block.EventsRoot(),
			"actual": block.WorldState().EventsRoot(),
		}).Debug("Failed to verify events.")
		return ErrInvalidBlockEventsRoot
	}

	// verify transaction root.
	if !reflect.DeepEqual(block.WorldState().ConsensusRoot(), block.ConsensusRoot()) {
		logging.VLog().WithFields(logrus.Fields{
			"expect": block.ConsensusRoot(),
			"actual": block.WorldState().ConsensusRoot(),
		}).Debug("Failed to verify dpos context.")
		return ErrInvalidBlockConsensusRoot
	}
	return nil
}

type verifyCtx struct {
	mergeCh chan bool
	block   *Block
}

// Execute block and return result.
func (block *Block) execute() error {
	startAt := time.Now().UnixNano()

	if err := block.rewardCoinbaseForMint(); err != nil {
		return err
	}

	context := &verifyCtx{
		mergeCh: make(chan bool, 1),
		block:   block,
	}
	dispatcher := dag.NewDispatcher(block.dependency, runtime.NumCPU(), 0, context, func(node *dag.Node, context interface{}) error { // TODO  time const   verify collect
		ctx := context.(*verifyCtx)
		block := ctx.block
		mergeCh := ctx.mergeCh
		tx := block.transactions[node.Index]
		metricsTxExecute.Mark(1)

		mergeCh <- true
		txWorldState, err := block.WorldState().Prepare(tx.Hash().String())
		if err != nil {
			<-mergeCh
			return err
		}
		<-mergeCh

		if _, err = block.ExecuteTransaction(tx, txWorldState); err != nil {
			return err
		}

		mergeCh <- true // TODO try to remove the lock
		if _, err := txWorldState.CheckAndUpdate(); err != nil {
			<-mergeCh
			return err
		}
		<-mergeCh

		return nil
	})

	start := time.Now().UnixNano()
	if err := dispatcher.Run(); err != nil {
		logging.CLog().Info("block verfiy txs err:", err, " dag: ", block.dependency.String())
		return err
	}
	end := time.Now().UnixNano()

	if len(block.transactions) != 0 {
		metricsTxVerifiedTime.Update((end - start) / int64(len(block.transactions)))
	} else {
		metricsTxVerifiedTime.Update(0)
	}

	if err := block.rewardCoinbaseForGas(); err != nil {
		return err
	}
	if err := block.WorldState().Flush(); err != nil {
		return err
	}

	endAt := time.Now().UnixNano()
	metricsBlockVerifiedTime.Update(endAt - startAt)
	metricsTxsInBlock.Update(int64(len(block.transactions)))

	return nil
}

// GetBalance returns balance for the given address on this block.
func (block *Block) GetBalance(address byteutils.Hash) (*util.Uint128, error) { // TODO return Account
	accState, err := block.WorldState().Clone()
	if err != nil {
		return nil, err
	}
	account, err := accState.GetOrCreateUserAccount(address)
	if err != nil {
		return nil, err
	}
	return account.Balance(), nil
}

// GetNonce returns nonce for the given address on this block.
func (block *Block) GetNonce(address byteutils.Hash) (uint64, error) {
	accState, err := block.WorldState().Clone()
	if err != nil {
		return 0, err
	}
	account, err := accState.GetOrCreateUserAccount(address)
	if err != nil {
		return 0, err
	}
	return account.Nonce(), nil
}

// FetchEvents fetch events by txHash.
func (block *Block) FetchEvents(txHash byteutils.Hash) ([]*state.Event, error) { // TODO clone first
	return block.WorldState().FetchEvents(txHash)
}

func (block *Block) rewardCoinbaseForMint() error {
	coinbaseAddr := block.Coinbase().Bytes()
	coinbaseAcc, err := block.WorldState().GetOrCreateUserAccount(coinbaseAddr)
	if err != nil {
		return err
	}
	logging.VLog().Info("rewardCoinbaseForMint ", "gas", BlockReward) // Refine: WithFields // TODO delete
	return coinbaseAcc.AddBalance(BlockReward)
}

func (block *Block) rewardCoinbaseForGas() error {
	worldState := block.WorldState()
	coinbaseAddr := (byteutils.Hash)(block.Coinbase().Bytes())

	logging.VLog().Info("rewardCoinbaseForGas") // TODO delete
	gasConsumed := worldState.GetGas()
	for from, gas := range gasConsumed {
		fromAddr, err := byteutils.FromHex(from)
		if err != nil {
			return err
		}
		logging.VLog().Info("rewardCoinbaseForGas from:", from, "gas", gas) // TODO delete

		if err := transfer(fromAddr, coinbaseAddr, gas, worldState); err != nil {
			return err
		}
	}
	return nil
}

// ExecuteTransaction execute the transaction
func (block *Block) ExecuteTransaction(tx *Transaction, ws WorldState) (bool, error) { // TODO system error: giveback, logic error: drop
	if giveback, err := CheckTransaction(tx, ws); err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"tx":  tx,
			"err": err,
		}).Info("Failed to check transaction")
		return giveback, err
	}

	if err := VerifyExecution(tx, block, ws); err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"tx":  tx,
			"err": err,
		}).Info("Failed to verify transaction execution")
		return false, err
	}

	if err := AcceptTransaction(tx, ws); err != nil {
		logging.VLog().WithFields(logrus.Fields{
			"tx":  tx,
			"err": err,
		}).Info("Failed to accept transaction")
		return false, err
	}

	return false, nil
}

// CheckContract check if contract is valid
func (block *Block) CheckContract(addr *Address) (state.Account, error) {

	worldState, err := block.worldState.Clone()
	if err != nil {
		return nil, err
	}
	return CheckContract(addr, worldState)
}

// GetTransaction from txs Trie
func (block *Block) GetTransaction(hash byteutils.Hash) (*Transaction, error) {
	worldState, err := block.worldState.Clone()
	if err != nil {
		return nil, err
	}
	return GetTransaction(hash, worldState)
}

// HashBlock return the hash of block.
func HashBlock(block *Block) (byteutils.Hash, error) { // TODO inter function
	hasher := sha3.New256()

	consensusRoot, err := proto.Marshal(block.ConsensusRoot())
	if err != nil {
		return nil, err
	}

	hasher.Write(block.ParentHash())
	hasher.Write(block.StateRoot())
	hasher.Write(block.TxsRoot())
	hasher.Write(block.EventsRoot())
	hasher.Write(consensusRoot)
	hasher.Write(block.header.coinbase.address)
	hasher.Write(byteutils.FromInt64(block.header.timestamp))
	hasher.Write(byteutils.FromUint32(block.header.chainID))

	for _, tx := range block.transactions {
		hasher.Write(tx.Hash())
	}

	return hasher.Sum(nil), nil
}

// HashPbBlock return the hash of pb block.
func HashPbBlock(pbBlock *corepb.Block) (byteutils.Hash, error) { // TODO nil check
	hasher := sha3.New256()

	consensusRoot, err := proto.Marshal(pbBlock.Header.ConsensusRoot)
	if err != nil {
		return nil, err
	}

	hasher.Write(pbBlock.Header.ParentHash) // TODO check header isn't nil
	hasher.Write(pbBlock.Header.StateRoot)
	hasher.Write(pbBlock.Header.TxsRoot)
	hasher.Write(pbBlock.Header.EventsRoot)
	hasher.Write(consensusRoot)
	hasher.Write(pbBlock.Header.Coinbase)
	hasher.Write(byteutils.FromInt64(pbBlock.Header.Timestamp))
	hasher.Write(byteutils.FromUint32(pbBlock.Header.ChainId))

	for _, tx := range pbBlock.Transactions {
		hasher.Write(tx.Hash)
	}

	return hasher.Sum(nil), nil
}

// RecoverMiner return miner from block
func RecoverMiner(block *Block) (*Address, error) { // TODO move to core/crypto.go. same as Transaction.Verify
	signature, err := crypto.NewSignature(keystore.Algorithm(block.Alg()))
	if err != nil {
		return nil, err
	}
	pub, err := signature.RecoverPublic(block.Hash(), block.Signature())
	if err != nil {
		return nil, err
	}
	pubdata, err := pub.Encoded()
	if err != nil {
		return nil, err
	}
	addr, err := NewAddressFromPublicKey(pubdata)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

// LoadBlockFromStorage return a block from storage
func LoadBlockFromStorage(hash byteutils.Hash, chain *BlockChain) (*Block, error) {
	if chain == nil {
		return nil, ErrNilArgument
	}

	value, err := chain.storage.Get(hash)
	if err != nil {
		return nil, err
	}
	pbBlock := new(corepb.Block)
	block := new(Block)
	if err = proto.Unmarshal(value, pbBlock); err != nil {
		return nil, err
	}
	if err = block.FromProto(pbBlock); err != nil {
		return nil, err
	}
	block.worldState, err = state.NewWorldState(chain.ConsensusHandler(), chain.storage)
	if err != nil {
		return nil, err
	}
	if err := block.WorldState().LoadAccountsRoot(block.StateRoot()); err != nil {
		return nil, err
	}
	if err := block.WorldState().LoadTxsRoot(block.TxsRoot()); err != nil {
		return nil, err
	}
	if err := block.WorldState().LoadEventsRoot(block.EventsRoot()); err != nil {
		return nil, err
	}
	if err := block.WorldState().LoadConsensusRoot(block.ConsensusRoot()); err != nil {
		return nil, err
	}
	block.sealed = true
	block.txPool = chain.txPool
	block.eventEmitter = chain.eventEmitter
	block.nvm = chain.nvm
	block.storage = chain.storage
	return block, nil
}

// Dispose dispose block.
func (block *Block) Dispose() {
	// cut off the parent block reference, prevent memory leak.
	block.parentBlock = nil
}
