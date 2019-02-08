test:
	go test -v -covermode=count -coverprofile=coverage.out ./...

lint:
	golangci-lint run
