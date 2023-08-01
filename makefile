lint:
	golangci-lint run -vvv


coverage:
	go test -v -coverprofile=coverage.txt -covermode=atomic ./...
