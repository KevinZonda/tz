.DEFAULT_GOAL := build

PORT := 8080

build:
	go build -v -o tz main.go

run:
	go run main.go ":$(PORT)"

pull:
	git pull

all: pull build run

.PHONY: build run