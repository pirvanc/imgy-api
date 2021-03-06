package main

import (
	"context"
	"imgy-api/image-services/crud"
	"imgy-api/pkg/dbaccessor"
	"imgy-api/pkg/httphandler"
	"imgy-api/pkg/logger"
	"imgy-api/pkg/validator"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is run by main
func Handler(context context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// create new pkg
	http := httphandler.NewHTTPHandler()
	logr := logger.NewLogger()
	vldr := validator.NewValidator()

	// get AWS Lambda environment variables
	awsRegion := string(os.Getenv("AWS_REGION_IMGY"))
	imagesTable := string(os.Getenv("IMAGES_TABLE"))
	bucketName := string(os.Getenv("S3_BUCKET_NAME"))

	// validate environment variables
	if vldr.IsEmptyString(awsRegion) == true || vldr.IsEmptyString(imagesTable) == true || vldr.IsEmptyString(bucketName) == true {
		logr.WriteLogEntry(nil, logr.GetErrorLogSeverity(), logr.GetEnvVarLogCode())

		return http.InternalServerError(), nil
	}

	// create new service
	dbac := dbaccessor.NewDBAccessor(awsRegion, imagesTable)
	crudService := crud.NewService(dbac, http, logr, vldr, awsRegion, bucketName)

	// obtain url path params
	imageID := crudService.GetImageIDPathParam(event)
	if imageID == nil {
		logr.WriteLogEntry(nil, logr.GetErrorLogSeverity(), logr.GetURLPathParamLogCode())

		return http.BadRequest(http.GetMissingPathParamErrorCode()), nil
	}

	// get response from service operation
	response := crudService.Get(*imageID)

	return response, nil
}

func main() {
	lambda.Start(Handler)
}
