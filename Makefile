.PHONY: run build

run: 
	@go run cmd/xeno/main.go

build: 
	@GOOS=linux GOARCH=arm64 go build -o xeno_arm ./cmd/xeno

build_all:
	echo "win-x64"
	@GOOS=windows GOARCH=amd64 go build -o bin/xeno-win-x64.exe ./cmd/xeno
	echo "win-x32"
	@GOOS=windows GOARCH=386 go build -o bin/xeno-win-x32.exe ./cmd/xeno

	echo "linux-x64"
	@GOOS=linux GOARCH=amd64 go build -o bin/xeno-linux-x64.bin ./cmd/xeno
	echo "linux-x32"
	@GOOS=linux GOARCH=386 go build -o bin/xeno-linux-x32.bin ./cmd/xeno

	echo "linux-arm64"
	@GOOS=linux GOARCH=arm64 go build -o bin/xeno-linux-arm64.bin ./cmd/xeno
	echo "linux-arm32"
	@GOOS=linux GOARCH=arm go build -o bin/xeno-linux-arm32.bin ./cmd/xeno