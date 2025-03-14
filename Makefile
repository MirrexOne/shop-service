.PHONY: build run
build:
	go build -v ./cmd/merch-shop

run:
	go run ./cmd/merch-shop/main.go
