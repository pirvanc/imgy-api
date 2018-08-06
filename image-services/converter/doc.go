package converter

import "github.com/aws/aws-lambda-go/events"

// DBAccessor interface
type DBAccessor interface {
	// CheckAuthorization(hashKeyName string, hashKeyValue string, rangeKeyName string, rangeKeyValue string) (interface{}, error)
	PutItem(item interface{}) (interface{}, error)
	GetItemByHashKey(hashKeyName string, hashKeyValue string) (interface{}, error)
	GetItemByHashKeyAndRangeKey(hashKeyName string, hashKeyValue string, hashRangeName string, rangeKeyValue string) (interface{}, error)
	DeleteItemByHashKey(hashKeyName string, hashKeyValue string) (interface{}, error)
}

// Logger interface
type Logger interface {
	// code
	GetInternalErrorLogCode() string
	GetJSONUnmarshalLogCode() string
	GetEnvVarLogCode() string
	GetAuthorizerClaimsLogCode() string
	GetURLPathParamLogCode() string
	GetAuthorizationLogCode() string
	GetDBCreateItemLogCode() string
	GetDBGetItemLogCode() string
	GetDBUpdateItemLogCode() string
	GetDBDeleteItemLogCode() string
	GetDBNotFoundItemLogCode() string
	GetDBExistingItemLogCode() string
	GetDBAttributeValueLogCode() string
	GetDBAttributeUnmarshalLogCode() string

	// severity
	GetInfoLogSeverity() string
	GetWarningLogSeverity() string
	GetErrorLogSeverity() string

	// write entry
	WriteLogEntry(message *string, severity string, code string)
}

// Validator interface
type Validator interface {
	IsEmptyString(string) bool
}

// HTTPHandler interface
type HTTPHandler interface {
	// error code
	GetDBItemNotFoundErrorCode() string
	GetDBCannotGetItemErrorCode() string
	GetDBCannotUpdateItemErrorCode() string
	GetDBCannotCreateItemErrorCode() string
	GetDBCannotDeleteItemErrorCode() string
	GetCannotDecodePayloadErrorCode() string
	GetMissingPathParamErrorCode() string
	GetDeletedItemErrorCode() string
	// response
	UnauthorizedRequest() events.APIGatewayProxyResponse
	ForbiddenRequest() events.APIGatewayProxyResponse
	InternalServerError() events.APIGatewayProxyResponse
	SuccessfulRequest(body interface{}) events.APIGatewayProxyResponse
	BadRequest(message string) events.APIGatewayProxyResponse
	NotFound(message string) events.APIGatewayProxyResponse
}
