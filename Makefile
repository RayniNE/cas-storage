build:
	@go build -o bin/output cmd/cas-storage/main.go

run: build
	 @./bin/output

test: 
	@go test ./... -v

coverage:
	@go test --coverfunc=c.out
	@go tool cover -func=c.out