package crud

import (
	"imgy-api/pkg/randomprovider"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GenerateURLforGetOperation is
func (s Service) GenerateURLforGetOperation() events.APIGatewayProxyResponse {

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
	return s.httphandler.SuccessfulRequest(response)
}

// GenerateURLforPutOperation is
func (s Service) GenerateURLforPutOperation() events.APIGatewayProxyResponse {

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
	return s.httphandler.SuccessfulRequest(response)
}
