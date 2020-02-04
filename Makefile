build:
	go build -o bin/go-logs-aggs cmd/main.go

deps:
	@cd cmd; go get -d -v; cd ..
