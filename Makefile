GO_CMD=go
GO_BUILD=$(GO_CMD) build
BUILD_DIR=./build

MAIN=cmd/main.go

BINARY_NAME=python_packer

all:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_linux_amd64 $(MAIN)
	GOOS=freebsd GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_freebsd_amd64 $(MAIN)
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_windows_amd64.exe $(MAIN)
