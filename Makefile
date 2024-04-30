ifeq ($(OS),Windows_NT)
    SOURCES := $(shell dir /S /B *.go)
else
    SOURCES := $(shell find . -name '*.go')
endif

ifeq ($(shell uname),Darwin)
    GOOS = darwin
    GOARCH = amd64
    EXEEXT =
else ifeq ($(shell uname),Linux)
    GOOS = linux
    GOARCH = amd64
    EXEEXT =
else ifeq ($(OS),Windows_NT)
    GOOS = windows
    GOARCH = amd64
    EXEEXT = .exe
endif

APP := howbob$(EXEEXT)
TARGET := ./dist/howbob_$(GOOS)_$(GOARCH)_v1/$(APP)

$(APP): $(TARGET)
	cp $< $@

$(TARGET): $(SOURCES)
	gofumpt -w $(SOURCES)
	golangci-lint run
	goreleaser build --single-target --snapshot --clean
	go vet ./...

all:
	goreleaser build --snapshot --clean

.PHONY: clean
clean:
	rm -f howbob
	rm -f $(TARGET)
	rm -rf dist
