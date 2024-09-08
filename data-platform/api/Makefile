BUILDPATH=$(CURDIR)/cmd

all: build run

build-target:
	@echo "start building tree..."

build: build-target
	@echo "start build..."
	@echo "${CURDIR}"
	@go build -o api ${BUILDPATH}/server.go

run:
	@echo "start running..."
	@./api

test:
	@echo "running tests..."
	@go test ./... -v
