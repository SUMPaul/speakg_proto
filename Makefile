all: stream_proxy


stream_proxy:
	export PATH=.:$PATH
	protoc --go_out=paths=source_relative:protobuf3.pb --proto_path=./protobuf3/ --micro_out=lang=go,plugins=grpc,paths=source_relative:protobuf3.pb stream_proxy/chat.proto