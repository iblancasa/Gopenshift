
VERSION ?= v0.0.1
BIN_DIR ?= bin
LD_FLAGS ?=
GO_FLAGS ?= CGO_ENABLED=0 GO111MODULE=on

# Run go fmt against code
.PHONY: fmt
fmt:
	go fmt ./...

# Run go vet against code
.PHONY: vet
vet:
	go vet ./...

# Run go lint against code
.PHONY: lint
lint:
	golangci-lint run ./...

# Build the is-openshift binary
.PHONY: build
build-is-openshift: fmt
	${GO_FLAGS} go build -ldflags $(LD_FLAGS) -o $(BIN_DIR)/is-openshift cmd/is-openshift/main.go

