build:
	go get github.com/aws/aws-lambda-go/lambda

	env GOOS=linux go build -ldflags="-s -w" -o bin/api-create cmd/image/create/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/api-get cmd/image/get/main.go
