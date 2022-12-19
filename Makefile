install:
	go get
	go install github.com/vektra/mockery/v2@latest
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	go mod tidy

run:
	go run main.go

tests:
	go clean -testcache
	go test -v ./...

mock-dependencies:
	mockery --keeptree --output mocks/service --dir service --all
	mockery --keeptree --output mocks/repository --dir repository --all

swagger-docs:
	swagger generate spec -m -o ./swagger.json
