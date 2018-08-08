# imgy-api

Image storage and processing service

## Getting Started

These instructions will get you a copy of the project for development on your local machine. See deployment for notes on how to deploy the project on a live system (AWS account is needed). 

### Prerequisites

Install serverless framework:

```
npm install -g serverless
```

## Running the tests

Unit test coverage: 100%

Run tests:

```
cd ~/imgy-api/cmd/image/{crud_operation}

go test
```

## Logging

CloudWatch is used for logging.

## Deployment

For deployment using AWS cloud formation, use the following commands:

```
make

serverless deploy --verbose --force
```

## Built With

* [Serverless](https://serverless.com/framework/docs/) - The serverless framework used to orchestrate and manage the lambdas
* [Golang](https://golang.org/doc/) - Programming language
* [DynamoDB](https://aws.amazon.com/documentation/dynamodb/) - Used for the peristent (database) layer
* [Cognito](https://aws.amazon.com/documentation/cognito/) - Used for the authentication layer 

## Authors

* **Cristian Pirvan** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

