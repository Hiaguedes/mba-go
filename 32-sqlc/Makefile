up:	
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
down:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
migrate:
	migrate create -ext sql -dir sql/migrations -seq $(name)
sqlc:
	sqlc generate
main: 
	go run cmd/run/main.go 
mainTx: 
	go run cmd/runTx/main.go
.PHONY: up down migrate