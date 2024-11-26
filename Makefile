BINARY := cvc

.PHONY: all clean test build lint fmt air

all: build

clean:
	rm -rf ./bin/*

build:
	mkdir -p bin
	go build -o bin/${BINARY} cmd/main.go

build-release:
	mkdir -p bin
	go build -o bin/${BINARY} -ldflags="-s -w" -trimpath cmd/main.go

run: build
	./bin/${BINARY}

test:
	go test ./...

lint:
	staticcheck ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

install: build-release
	sudo cp ./bin/${BINARY} /usr/local/bin/${BINARY}

watch:
	watchexec "make lint"
