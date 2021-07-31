build:
	go build -o bin/registry cmd/registry/main.go

build-all:
	echo "Compiling for every OS and Platform"
	GOOS=windows GOARCH=amd64 go build -o bin/registry-windows-amd64 cmd/registry/main.go
	GOOS=linux GOARCH=arm go build -o bin/registry-linux-arm cmd/registry/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/registry-linux-arm64 cmd/registry/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/registry-freebsd-386 cmd/registry/main.go
