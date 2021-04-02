# -- Variables --
VERSION = $(shell cat VERSION.txt)
TAG ?= ${VERSION}
IMAGE = "twitapp:${TAG}"
# --

.PHONY: image
image:
	$(call yello, "Building Docker image...")
	docker build . -t ${IMAGE}
	$(call yello, "Docker image built succesfully")

.PHONY: lint
lint:
	$(call yellow, "Linting...")
	@golangci-lint run -c .golanci.yml
	$(call yellow, "Linting completed...")

.PHONY: build_mac
build_mac: build_modules
	$(call yellow, "Building binaries for MacOS...")
	$(call build_packages,darwin)
	$(call green, "Binaries built succesfully...")

.PHONY: build_linux
build_linux: build_modules
	$(call yellow, "Building binaries for Linux...")
	$(call build_packages, "linux")
	$(call green, "Binaries built succesfully...")

build_modules:
	$(call yellow, "Building modules...")
	@go mod vendor && go mod tidy
	$(call green, "Modules built succesfully...")

# -- Functions --
define build_packages
	@mkdir -p bin
	@GOOS=$1 GOARCH=386 go build -o bin/twix-${VERSION}.$1.386 cmd/main.go > /dev/null 2>&1 || true
	@GOOS=$1 GOARCH=amd64 go build -o bin/twix-${VERSION}.$1.amd64 cmd/main.go  > /dev/null 2>&1 || true
	@GOOS=$1 GOARCH=arm go build -o bin/twix-${VERSION}.$1.arm cmd/main.go  > /dev/null 2>&1 || true
endef

define yellow
	@tput setaf 3
	@echo $1
	@tput sgr0
endef

define green
	@tput setaf 2
	@echo $1
	@tput sgr0
endef

