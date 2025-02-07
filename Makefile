MODULE_NAME = $(shell go list -m)
VERSION = $(shell git describe --tags --abbrev=0 2>/dev/null || echo "")

SOURCE_DIR = ./src/cmd/app
BUILD_DIR = ./build
BINARY_NAME = bin

LDFLAGS += -X "$(MODULE_NAME)/src//pkg/motd.Version=$(VERSION)"
LDFLAGS += -X "$(MODULE_NAME)/src/pkg/motd.Module=$(MODULE_NAME)"

build:
	go build -ldflags='$(LDFLAGS)' -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)
	
run:
	./build/bin

clean:
	rm -rf $(BUILD_DIR)