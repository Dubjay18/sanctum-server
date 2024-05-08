
build:
	   @go build -o bin/sanctum
run: build
	   @./bin/sanctum
dev: build
	   nodemon --exec go run main.go --signal SIGTERM

test:
	   @go test -v ./...