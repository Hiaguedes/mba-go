.PHONY: run
run:
	GODEBUG=gctrace=1 go run main.go
run-300:
	GODEBUG=gctrace=1 GOGC=300 go run main.go
disable-gc:
	GODEBUG=gctrace=1 GOGC=-1 go run main.go