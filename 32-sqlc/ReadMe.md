# Comandos usados

migrate create -ext=sql -dir=sql/migrations -seq init 

migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up

migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down