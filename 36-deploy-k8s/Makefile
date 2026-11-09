.PHONY: run
run: 
	go run main.go
build:
	docker build -t mba-go:latest .
build-prod:
	docker build -f Dockerfile.prod -t hiaguedes/build-k8s:prod .