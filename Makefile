protoc:
	protoc --go_out=./ \
    --go-grpc_out=./ \
    ./api/proto/*.proto
grpc-protoc:
	protoc --go_out=./ \
	--go-grpc_out=./ \
	./api/proto/messages/*.proto
msg-protoc:
	protoc --go_out=./ \
    --go-grpc_out=./ \
    ./internal/infrastructure/messagebus/messages/*.proto
