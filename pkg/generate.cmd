echo "GENERATE PROTO"

protoc --proto_path=./ --go-micro_out=components="micro|server|client":./ --micro_out=./ --go_out=./ ./grpc.proto