package main

import (
	"imgy-api/image-services/crud"
	"imgy-api/pkg/dbaccessor"
	"imgy-api/pkg/httphandler"
	"imgy-api/pkg/logger"
	"imgy-api/pkg/validator"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ConfigurationFile = "settings.json"
	AwsRegion         = "eu-west-1"
	ImagesTable       = "staging-images"
	BucketName        = "staging-imageservice-bucket"
	ImageID           = "bdk183ciotqg01a9arng"
)

func TestImageCreate(t *testing.T) {

	var err error

	// get values from settings.json
	configFile, err := os.Open(ConfigurationFile)
	defer configFile.Close()
	assert.Nil(t, err)

	// create new pkg
	http := httphandler.NewHTTPHandler()
	logr := logger.NewLogger()
	vldr := validator.NewValidator()

	// get response from service operation
	dbac := dbaccessor.NewDBAccessor(AwsRegion, ImagesTable)
	crudService := crud.NewService(dbac, http, logr, vldr, AwsRegion, BucketName)
	response := crudService.Get(ImageID)
	assert.Equal(t, response.StatusCode, 200)
}
