package crud

import (
	"encoding/json"
	"imgy-api/pkg/randomprovider"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rs/xid"
)

// Response is
type Response struct {
	ImageID   string `json:"imageID"`
	SignedURL string `json:"signedURL"`
}

// Service is...
type Service struct {
	dbaccessor  DBAccessor
	httphandler HTTPHandler
	logger      Logger
	validator   Validator
	awsregion   string
	bucketname  string
}

// NewService is...
func NewService(dbac DBAccessor, http HTTPHandler, logg Logger, vald Validator, awsregion string, bucketname string) Service {
	return Service{
		dbaccessor:  dbac,
		httphandler: http,
		logger:      logg,
		validator:   vald,
		awsregion:   awsregion,
		bucketname:  bucketname,
	}
}

// ImageIDPathParam is
const (
	ImageIDPathParam string = "imageID"
	ImageHashKeyName string = "imageID"
)

// Create is
func (s Service) Create(newImage NewImage) events.APIGatewayProxyResponse {

	// generate new id
	imageID := xid.New().String()

	// create new object
	var image Image
	image.ImageID = imageID

	// map fields from new object
	image.Description = newImage.Description

	// fill in default values
	image.CreatedAt = time.Now().Format(time.RFC3339)
	image.CreatedAtTimestamp = time.Now().Unix()

	// create object
	_, err := s.dbaccessor.PutItem(image)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetDBCreateItemLogCode())

		return s.httphandler.InternalServerError()
	}

	// create new aws session
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(s.awsregion)},
	)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetInternalErrorLogCode())

		return s.httphandler.InternalServerError()
	}

	// create S3 service client
	svc := s3.New(session)
	uploadKey, err := randomprovider.RandomHex(20)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetInternalErrorLogCode())

		return s.httphandler.InternalServerError()
	}

	// get signed url
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucketname),
		Key:         aws.String(uploadKey),
		ContentType: aws.String("image"),
		ACL:         aws.String("public-read"),
	})
	signedURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetInternalErrorLogCode())

		return s.httphandler.InternalServerError()
	}

	// return successful response
	var response Response
	response.SignedURL = signedURL
	response.ImageID = imageID
	return s.httphandler.SuccessfulRequest(response)
}

// Get is
func (s Service) Get(imageID string) events.APIGatewayProxyResponse {

	// get object
	image := Image{}
	resultItem, err := s.dbaccessor.GetItemByHashKey(ImageHashKeyName, imageID)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetDBGetItemLogCode())

		return s.httphandler.NotFound("Image not found")
	}
	assertedResultItem, ok := resultItem.(map[string]*dynamodb.AttributeValue)
	if !ok {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetDBAttributeValueLogCode())

		return s.httphandler.NotFound("Image not found")
	}
	err = dynamodbattribute.UnmarshalMap(assertedResultItem, &image)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetJSONUnmarshalLogCode())

		return s.httphandler.NotFound("Image not found")
	}
	if resultItem == nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetDBNotFoundItemLogCode())

		return s.httphandler.NotFound("Image not found")
	}

	// create new aws session
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(s.awsregion)},
	)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetInternalErrorLogCode())

		return s.httphandler.InternalServerError()
	}

	// create S3 service client
	svc := s3.New(session)
	uploadKey, err := randomprovider.RandomHex(20)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetInternalErrorLogCode())

		return s.httphandler.InternalServerError()
	}

	// get signed url
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucketname),
		Key:    aws.String(uploadKey),
	})
	signedURL, err := req.Presign(15 * time.Minute)
	if err != nil {
		errMessage := err.Error()
		s.logger.WriteLogEntry(&errMessage, s.logger.GetErrorLogSeverity(), s.logger.GetInternalErrorLogCode())

		return s.httphandler.InternalServerError()
	}

	// return successful response
	var response Response
	response.SignedURL = signedURL
	response.ImageID = imageID
	return s.httphandler.SuccessfulRequest(response)
}

// DecodePostPayload is
func (s Service) DecodePostPayload(event events.APIGatewayProxyRequest) (*NewImage, error) {

	var newImage = &NewImage{}
	bodyBytes := []byte(event.Body)
	err := json.Unmarshal(bodyBytes, newImage)
	if err != nil {

		return nil, err
	}

	return newImage, nil
}

// GetImageIDPathParam is
func (s Service) GetImageIDPathParam(event events.APIGatewayProxyRequest) *string {

	// try to get path param
	imageIDPathParam := event.PathParameters[ImageIDPathParam]
	if len(imageIDPathParam) == 0 {

		return nil
	}

	return &imageIDPathParam
}
