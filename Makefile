all: stream_proxy


stream_proxy:
	export PATH=.:$PATH
	protoc --go_out=paths=source_relative:protobuf3.pb --proto_path=./protobuf3/ --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:protobuf3.pb stream_proxy/chat.proto