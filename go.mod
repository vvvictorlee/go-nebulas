module github.com/nebulasio/go-nebulas

go 1.12

require (
	cloud.google.com/go v0.46.3 // indirect
	github.com/VividCortex/godaemon v0.0.0-20150910212227-3d9f6e0b234f
	github.com/btcsuite/btcd v0.0.0-20190523000118-16327141da8c
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/gogo/protobuf v1.3.0
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/pprof v0.0.0-20190908185732-236ed259b199 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.2
	github.com/hashicorp/golang-lru v0.5.3
	github.com/influxdata/influxdb v1.7.8
	github.com/kr/pty v1.1.8 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20190725011945-5c849dd2c51d // indirect
	github.com/libp2p/go-libp2p v0.3.1
	github.com/libp2p/go-libp2p-core v0.2.2
	github.com/libp2p/go-libp2p-kbucket v0.2.1
	github.com/libp2p/go-libp2p-peerstore v0.1.3
	github.com/libp2p/go-libp2p-swarm v0.2.1
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/multiformats/go-multicodec v0.1.6
	github.com/peterh/liner v1.1.0
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/rogpeppe/go-internal v1.3.2 // indirect
	github.com/rs/cors v1.7.0
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tecbot/gorocksdb v0.0.0-20190705090504-162552197222
	github.com/urfave/cli v1.22.1
	github.com/willf/bitset v1.1.10 // indirect
	github.com/willf/bloom v2.0.3+incompatible
	go.opencensus.io v0.22.1 // indirect
	golang.org/x/crypto v0.0.0-20190923035154-9ee001bba392
	golang.org/x/exp v0.0.0-20190919035709-81c71964d733 // indirect
	golang.org/x/image v0.0.0-20190910094157-69e4b8554b2a // indirect
	golang.org/x/mobile v0.0.0-20190910184405-b558ed863381 // indirect
	golang.org/x/net v0.0.0-20190921015927-1a5e07d1ff72
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/time v0.0.0-20190921001708-c4c64cad1fd0 // indirect
	golang.org/x/tools v0.0.0-20190920225731-5eefd052ad72 // indirect
	google.golang.org/api v0.10.0 // indirect
	google.golang.org/appengine v1.6.3 // indirect
	google.golang.org/genproto v0.0.0-20190916214212-f660b8655731
	google.golang.org/grpc v1.23.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/grpc-ecosystem/grpc-gateway => github.com/nebulasio/grpc-gateway v1.11.2
