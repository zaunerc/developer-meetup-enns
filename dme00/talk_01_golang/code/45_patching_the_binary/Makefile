BINARY=main

BUILD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-w -s -X main.GitCommitHash=${BUILD}"

build:
	env GOOS=linux   GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-linux-amd64
	env GOOS=freebsd GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}-freebsd-amd64
