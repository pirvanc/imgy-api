build:
	go get github.com/aws/aws-lambda-go/lambda
	env GOOS=linux go build -ldflags="-s -w" -o bin/api-uploader cmd/image/upload/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/api-downloader cmd/image/download/main.go
