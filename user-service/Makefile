build:
	protoc -I . --micro_out=. --go_out=. ./proto/user/user.proto \
		docker build -t user-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service