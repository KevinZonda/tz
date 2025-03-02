.DEFAULT_GOAL := build

build:
	go build -v -o tz main.go

run:
	go run main.go $(PORT)


.PHONY: build run