.PHONY: run build

run: 
	@go run cmd/xeno/main.go

build: 
	@GOOS=linux GOARCH=arm64 go build -o xeno_arm ./cmd/xeno