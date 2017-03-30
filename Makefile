default: build

bootstrap:
	go get github.com/urfave/cli
	brew install sdl2
	make updatedeps

build: clean vet lint
	GOOS=darwin GOARCH=amd64 godep go build -v -o ./bin/GChip8 ./src/main

clean:
	rm -rf ./bin/*

vet:
	go vet ./src/...

lint:
	golint ./src/...

updatedeps:
	rm -rf vendor
	rm lock.json
	dep init
	
