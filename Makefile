#Makefile

.PHONY: build run rundev test

build:
	go build -o calculator.exe

run:
	./calculator.exe

rundev:
	go run main.go

test:
	go test ./... -v

all: build test run