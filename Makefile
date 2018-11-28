.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/authenticate authenticate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/costIndex costIndex/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
