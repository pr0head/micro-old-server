echo "GENERATE PROTO"

protoc --proto_path=./ --go-micro_out=components="micro|server|client|rpc":./ --micro_out=./ --go_out=./ ./grpc.proto