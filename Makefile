all:
	go build cmd/leveldbctl/leveldbctl.go
test:
	go test -v -covermode=count -coverprofile=coverage.out ./...

lint:
	golangci-lint run
