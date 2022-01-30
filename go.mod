module github.com/pr0head/micro-old-server

go 1.13

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro v1.18.0
	go.unistack.org/micro-network-transport-http/v3 v3.8.0
	go.unistack.org/micro-register-etcd/v3 v3.8.1
	go.unistack.org/micro/v3 v3.8.20
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
