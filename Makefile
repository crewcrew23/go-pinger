.PHONY: build clean

BINARY_NAME := pinger
BIN_DIR := ./bin
MAIN_PKG := ./cmd/pinger/

# Определяем текущую платформу
ifeq ($(OS),Windows_NT)
    CURRENT_OS := windows
    BINARY_EXT := .exe
    SHELL := powershell.exe
    .SHELLFLAGS := -NoProfile -Command
    RM := Remove-Item -Force -ErrorAction SilentlyContinue
    MKDIR := New-Item -ItemType Directory -Force -ErrorAction SilentlyContinue
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        CURRENT_OS := linux
    endif
    ifeq ($(UNAME_S),Darwin)
        CURRENT_OS := darwin
    endif
    BINARY_EXT :=
    RM := rm -f
    MKDIR := mkdir -p
endif

build: $(BIN_DIR)
	@echo "Building for current platform ($(CURRENT_OS))..."
ifeq ($(OS),Windows_NT)
	@$$env:GOOS="$(CURRENT_OS)"; $$env:GOARCH="amd64"; go build -o "$(BIN_DIR)/$(BINARY_NAME)$(BINARY_EXT)" "$(MAIN_PKG)"
else
	@GOOS=$(CURRENT_OS) GOARCH=amd64 go build -o "$(BIN_DIR)/$(BINARY_NAME)$(BINARY_EXT)" "$(MAIN_PKG)"
endif
	@echo "Build complete. Binary created at $(BIN_DIR)/$(BINARY_NAME)$(BINARY_EXT)"

$(BIN_DIR):
	@$(MKDIR) $(BIN_DIR)

clean:
	@echo "Cleaning binaries..."
	@$(RM) $(BIN_DIR)/$(BINARY_NAME)*
	@echo "Clean complete"

.DEFAULT_GOAL := build