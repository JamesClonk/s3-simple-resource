all: built-check built-in built-out

built-check: check/main.go
	GOARCH=amd64 GOOS=linux go build -o built-check check/main.go

built-in: in/main.go
	GOARCH=amd64 GOOS=linux go build -o built-in in/main.go

built-out: out/main.go
	GOARCH=amd64 GOOS=linux go build -o built-out out/main.go
