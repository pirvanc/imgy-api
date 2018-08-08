# imgy-api

Image storage and processing service - REST API

The API consists of a series of FaaS (function-as-a-service) orchestrated by AWS API Gateway. 

Authentication was implemented using JWT tokens issued by AWS Cognito.

## Getting Started

These instructions will get you a copy of the project for development on your local machine. For deployment on a live system an AWS account is needed. 


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

* [Serverless](https://serverless.com/framework/docs/) - The serverless framework used to manage the lambdas
* [Golang](https://golang.org/doc/) - Programming language
* [DynamoDB](https://aws.amazon.com/documentation/dynamodb/) - Used for the peristent (database) layer
* [Cognito](https://aws.amazon.com/documentation/cognito/) - Used for the authentication layer (JWT)
* [API Gateway](https://aws.amazon.com/documentation/apigateway/) - Used for orchestration of lambdas
* [IAM](https://aws.amazon.com/documentation/iam/) - Used for identity and access management

## Authors

* **Cristian Pirvan** - *Initial work* - [pirvanc](https://github.com/pirvanc)


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

