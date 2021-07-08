.PHONY: build
build: deps
	./build.sh

.PHONY: build_linux
build_linux: deps
	./build_linux.sh

.PHONY: test
test: deps
	./test.sh

.PHONY: deps
deps:
	go get -t -v ./...

.PHONY: image
image:
	docker build -f Dockerfile -t lal:latest .

.PHONY: clean
clean:
	rm -rf ./bin ./release ./logs

.PHONY: all
all: build test
