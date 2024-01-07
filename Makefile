protoc:
	protoc --go_out=./ \
    --go-grpc_out=./ \
    ./api/proto/*.proto

msg-protoc:
	protoc --go_out=./ \
    --go-grpc_out=./ \
    ./internal/infrastructure/messagebus/messages/*.proto
