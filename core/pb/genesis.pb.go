// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: genesis.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	genesis.proto

It has these top-level messages:
	Genesis
	GenesisMeta
	GenesisConsensus
	GenesisConsensusDpos
	GenesisTokenDistribution
*/
package corepb

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Genesis struct {
	// genesis meta
	Meta *GenesisMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	// genesis consensus config
	Consensus *GenesisConsensus `protobuf:"bytes,2,opt,name=consensus" json:"consensus,omitempty"`
	// genesis token distribution address
	// map<string, string> token_distribution = 3;
	TokenDistribution []*GenesisTokenDistribution `protobuf:"bytes,3,rep,name=token_distribution,json=tokenDistribution" json:"token_distribution,omitempty"`
}

func (m *Genesis) Reset()                    { *m = Genesis{} }
func (m *Genesis) String() string            { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()               {}
func (*Genesis) Descriptor() ([]byte, []int) { return fileDescriptorGenesis, []int{0} }

func (m *Genesis) GetMeta() *GenesisMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Genesis) GetConsensus() *GenesisConsensus {
	if m != nil {
		return m.Consensus
	}
	return nil
}

func (m *Genesis) GetTokenDistribution() []*GenesisTokenDistribution {
	if m != nil {
		return m.TokenDistribution
	}
	return nil
}

type GenesisMeta struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *GenesisMeta) Reset()                    { *m = GenesisMeta{} }
func (m *GenesisMeta) String() string            { return proto.CompactTextString(m) }
func (*GenesisMeta) ProtoMessage()               {}
func (*GenesisMeta) Descriptor() ([]byte, []int) { return fileDescriptorGenesis, []int{1} }

func (m *GenesisMeta) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

type GenesisConsensus struct {
	// ChainID.
	Dpos *GenesisConsensusDpos `protobuf:"bytes,1,opt,name=dpos" json:"dpos,omitempty"`
}

func (m *GenesisConsensus) Reset()                    { *m = GenesisConsensus{} }
func (m *GenesisConsensus) String() string            { return proto.CompactTextString(m) }
func (*GenesisConsensus) ProtoMessage()               {}
func (*GenesisConsensus) Descriptor() ([]byte, []int) { return fileDescriptorGenesis, []int{2} }

func (m *GenesisConsensus) GetDpos() *GenesisConsensusDpos {
	if m != nil {
		return m.Dpos
	}
	return nil
}

type GenesisConsensusDpos struct {
	// dpos genesis dynasty address
	Dynasty []string `protobuf:"bytes,1,rep,name=dynasty" json:"dynasty,omitempty"`
}

func (m *GenesisConsensusDpos) Reset()                    { *m = GenesisConsensusDpos{} }
func (m *GenesisConsensusDpos) String() string            { return proto.CompactTextString(m) }
func (*GenesisConsensusDpos) ProtoMessage()               {}
func (*GenesisConsensusDpos) Descriptor() ([]byte, []int) { return fileDescriptorGenesis, []int{3} }

func (m *GenesisConsensusDpos) GetDynasty() []string {
	if m != nil {
		return m.Dynasty
	}
	return nil
}

type GenesisTokenDistribution struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Value   string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GenesisTokenDistribution) Reset()                    { *m = GenesisTokenDistribution{} }
func (m *GenesisTokenDistribution) String() string            { return proto.CompactTextString(m) }
func (*GenesisTokenDistribution) ProtoMessage()               {}
func (*GenesisTokenDistribution) Descriptor() ([]byte, []int) { return fileDescriptorGenesis, []int{4} }

func (m *GenesisTokenDistribution) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GenesisTokenDistribution) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Genesis)(nil), "corepb.Genesis")
	proto.RegisterType((*GenesisMeta)(nil), "corepb.GenesisMeta")
	proto.RegisterType((*GenesisConsensus)(nil), "corepb.GenesisConsensus")
	proto.RegisterType((*GenesisConsensusDpos)(nil), "corepb.GenesisConsensusDpos")
	proto.RegisterType((*GenesisTokenDistribution)(nil), "corepb.GenesisTokenDistribution")
}

func init() { proto.RegisterFile("genesis.proto", fileDescriptorGenesis) }

var fileDescriptorGenesis = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xa9, 0x9d, 0xab, 0xbd, 0x65, 0xa0, 0x71, 0x0f, 0x11, 0x7c, 0x28, 0x79, 0xb1, 0x4f,
	0x65, 0x4c, 0xf0, 0x0f, 0x58, 0x10, 0x05, 0x11, 0x82, 0xef, 0x23, 0x6d, 0x2e, 0x1a, 0xd4, 0xa4,
	0xf4, 0xa6, 0xc2, 0x7e, 0x9b, 0x7f, 0x4e, 0x96, 0xae, 0x38, 0x8a, 0x7b, 0x3c, 0x39, 0x5f, 0x2e,
	0xe7, 0xdc, 0x0b, 0x8b, 0x37, 0xb4, 0x48, 0x86, 0xca, 0xb6, 0x73, 0xde, 0xb1, 0x79, 0xe3, 0x3a,
	0x6c, 0x6b, 0xf1, 0x13, 0x41, 0xf2, 0x30, 0x38, 0xec, 0x06, 0x66, 0x5f, 0xe8, 0x15, 0x8f, 0xf2,
	0xa8, 0xc8, 0xd6, 0x97, 0xe5, 0x80, 0x94, 0x7b, 0xfb, 0x19, 0xbd, 0x92, 0x01, 0x60, 0x77, 0x90,
	0x36, 0xce, 0x12, 0x5a, 0xea, 0x89, 0x9f, 0x04, 0x9a, 0x4f, 0xe8, 0xfb, 0xd1, 0x97, 0x7f, 0x28,
	0x7b, 0x01, 0xe6, 0xdd, 0x07, 0xda, 0x8d, 0x36, 0xe4, 0x3b, 0x53, 0xf7, 0xde, 0x38, 0xcb, 0xe3,
	0x3c, 0x2e, 0xb2, 0x75, 0x3e, 0x19, 0xf0, 0xba, 0x03, 0xab, 0x03, 0x4e, 0x5e, 0xf8, 0xe9, 0x93,
	0x28, 0x20, 0x3b, 0x48, 0xc7, 0xae, 0xe0, 0xac, 0x79, 0x57, 0xc6, 0x6e, 0x8c, 0x0e, 0x25, 0x16,
	0x32, 0x09, 0xfa, 0x51, 0x8b, 0x0a, 0xce, 0xa7, 0xc9, 0xd8, 0x0a, 0x66, 0xba, 0x75, 0xb4, 0xef,
	0x7b, 0x7d, 0xac, 0x41, 0xd5, 0x3a, 0x92, 0x81, 0x14, 0x2b, 0x58, 0xfe, 0xe7, 0x32, 0x0e, 0x89,
	0xde, 0x5a, 0x45, 0x7e, 0xcb, 0xa3, 0x3c, 0x2e, 0x52, 0x39, 0x4a, 0xf1, 0x04, 0xfc, 0x58, 0xa1,
	0xdd, 0x2f, 0xa5, 0x75, 0x87, 0x34, 0x44, 0x48, 0xe5, 0x28, 0xd9, 0x12, 0x4e, 0xbf, 0xd5, 0x67,
	0x8f, 0x61, 0xb9, 0xa9, 0x1c, 0x44, 0x3d, 0x0f, 0xa7, 0xbb, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0xf4, 0x0e, 0x3f, 0xcb, 0x01, 0x00, 0x00,
}
