build:
	go build -o bin/speed-bank

run: build
	./bin/speed-bank

test:
	go test -v ./...

tidy:
	go mod tidy