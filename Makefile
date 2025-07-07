BINARY_NAME=lucerna

build:
	go build -o ${BINARY_NAME} ./cmd/server

run:
	go run ./cmd/server/

clean:
	rm -f ${BINARY_NAME}
	rm -rf logs/
	go mod tidy	