.PHONY: gen-env gen-mock test build clean

gen-env:
	cp ./config/.env.default .env
	
gen-mock:
	go generate ./...

test:
	go test ./... -cover

build:
	go mod tidy
	go build -o build/api_server cmd/api/main.go
	
clean:
	rm -rf build