package main

import (
	"imgy-api/image-services/urlsigner"
	"imgy-api/pkg/dbaccessor"
	"imgy-api/pkg/httphandler"
	"imgy-api/pkg/logger"
	"imgy-api/pkg/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateURLforGetOperation(t *testing.T) {

	// create new pkg
	http := httphandler.NewHTTPHandler()
	logr := logger.NewLogger()
	vldr := validator.NewValidator()

	// get AWS Lambda environment variables
	awsRegion := "eu-west-1"
	imagesTable := "dev-images"
	bucketName := "imgy"

	// create new service
	dbac := dbaccessor.NewDBAccessor(awsRegion, imagesTable)
	urlSignerService := urlsigner.NewService(dbac, http, logr, vldr, awsRegion, bucketName)

	// get response from service operation
	response := urlSignerService.GenerateURLforGetOperation()

	// assert
	assert.EqualValues(t, 200, response.StatusCode)
}
