protocgen:
	protoc --go_out=. -I ./api/proto/ ./api/proto/*.proto --go-grpc_out=.
microservice-http:
	go run cmd/http/main.go
microservice-grpc:
	go run cmd/grpc/main.go