BINARY_NAME=laterna

build:
	mkdir -p build
	cp config.json build
	go build -o ./build/${BINARY_NAME} ./cmd/server

run:
	go run ./cmd/server/

clean:
	rm -rf build/
	rm -rf logs/
	go mod tidy	