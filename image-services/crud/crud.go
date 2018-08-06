package crud

import (
	"time"

	"github.com/rs/xid"
)

// Response is
type Response struct {
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

// CreateImage is
func (s Service) CreateImage(newImage NewImage) (*string, bool) {

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

		return nil, false
	}

	return &imageID, true
}
