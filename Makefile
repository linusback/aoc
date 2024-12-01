ENTRYPOINT=./main.go
BINARY_NAME=aoc2024

BENCH=

DAY=1

run: build
	./dist/$(BINARY_NAME) $(DAY)

dev:
	go run $(ENTRYPOINT_NAME) $(DAY)

download:
	

all: tests build

# perhaps add GOAMD64=v3 to architecture
build:
	go build -ldflags='-s -w' -o dist/$(BINARY_NAME) $(ENTRYPOINT_NAME)

test-out-pipe: build
	./dist/${BINARY_NAME2} --pipe-to="$(ENTRYPOINT_NAME3) ./tmp/test_file.txt"

tests:
	go test ./...

bench:
	go test ./... -bench=. -benchtime 3s -run=^\# -cpu=1,20

bench-prof:
	go test . -bench=${BENCH} -benchtime 3s -run=^\# -cpu=20 -cpuprofile ./tmp/$(subst /,-,$(BENCH))_cpu.prof -memprofile ./tmp/$(subst /,-,$(BENCH))_mem.prof -o ./tmp/$(subst /,-,$(BENCH)).test


clean:
	rm -f dist/*
	go mod tidy
	go clean

vet:
	go vet

lint:
	golangci-lint run --enable-all
