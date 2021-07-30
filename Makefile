build:
	go build -o bin/file-server cmd/file-server/main.go

build-all:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/file-server-linux-arm cmd/file-server/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/file-server-linux-arm64 cmd/file-server/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/file-server-freebsd-386 cmd/file-server/main.go
