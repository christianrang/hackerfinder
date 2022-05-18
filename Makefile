.PHONY:build
build:
	mkdir -p bin
	GOOS=linux GOARCH=386 go build -o bin/hackerfinder_linux_386 -v cmd/cli/main.go
	GOOS=linux GOARCH=arm go build -o bin/hackerfinder_linux_arm -v cmd/cli/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/hackerfinder_linux_arm64 -v cmd/cli/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/hackerfinder_linux_amd64 -v cmd/cli/main.go
