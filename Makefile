fmt :
	go fmt ./...

test : fmt
	go test ./... -bench .

build : fmt
	go build -v -o ./build/ ./...
