MAIN_PKG=./cmd/server

BIN_NAME ?= laterna
BIN_DIR=build


build:
	@mkdir -p ${BIN_DIR}
	@echo "🔧 Building binary using the standard go compiler..."
	cp config.json build
	go build -o $(BIN_DIR)/$(BIN_NAME) $(MAIN_PKG)
	@echo "✅ Build complete: $(BIN_DIR)/$(BIN_NAME)"
	
run:
	go run ./cmd/server/
clean:
	rm -rf ${BIN_DIR}
	rm -rf logs/
	go mod tidy	
