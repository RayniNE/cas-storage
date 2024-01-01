build:
	@go build -o bin/output cmd/cas-storage/main.go

run: build
	 @./bin/output

test: 
	@go test ./... -v

coverage: test
	@go test --coverprofile=c.out ./peer2peer
	@go tool cover -func=c.out