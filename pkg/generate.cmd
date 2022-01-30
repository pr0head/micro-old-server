echo "GENERATE PROTO"

protoc --proto_path=./ --go-micro_out=components="micro|server|client|http":./ --micro_out=./ --go_out=./ ./grpc.proto
