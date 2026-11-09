protogen:
	protoc --go_out=. --go-grpc_out=. proto/*.proto
server:
	go run cmd/grpc/main.go
evans:
	~/go/bin/evans --proto proto/*.proto repl