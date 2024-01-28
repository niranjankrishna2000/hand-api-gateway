proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/**/pb/*.proto

server:
	go run cmd/main.go

swag: ## Generate swagger docs
	swag fmt
	swag init -g ./pkg/**/routes/*.go -o ./cmd/
