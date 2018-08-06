package httphandler

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// HTTPHandler is...
type HTTPHandler struct {
	errorCode errorCode
}

// NewHTTPHandler is...
func NewHTTPHandler() HTTPHandler {

	defaultErrorCode := errorCode{
		DBItemNotFound:      "Item not found in database",
		DBCannotCreateItem:  "Item cannot be created",
		DBCannotGetItem:     "Item cannot be retrieved",
		DBCannotUpdateItem:  "Item cannot be updated",
		DBCannotDeleteItem:  "Item cannot be deleted",
		CannotDecodePayload: "Cannot decode payload",
		MissingPathParam:    "Cannot obtain path parameter",
		DeletedItem:         "Item was deleted",
	}

	return HTTPHandler{
		errorCode: defaultErrorCode,
	}
}

// responseBody is...
type responseBody struct {
	Message interface{} `json:"message"`
}

// GetUserEmailFromAuthorizerClaims is ..
func (h HTTPHandler) GetUserEmailFromAuthorizerClaims(event events.APIGatewayProxyRequest) *string {

	// get claims
	claims := event.RequestContext.Authorizer["claims"]
	if claims == nil {
		return nil
	}
	claimsDetails, ok := claims.(map[string]interface{})
	if !ok {
		return nil
	}
	// get email
	accountUserID, ok := claimsDetails["email"].(string)
	if !ok {
		return nil
	}

	return &accountUserID
}

// UnauthorizedRequest 401
func (h HTTPHandler) UnauthorizedRequest() events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	var body responseBody
	body.Message = "Request is unauthorized"
	response.StatusCode = 401

	content, _ := json.Marshal(body)
	response.Body = string(content)

	return response
}

// ForbiddenRequest 403
func (h HTTPHandler) ForbiddenRequest() events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	var body responseBody
	body.Message = "Request is forbidden"
	response.StatusCode = 403

	content, _ := json.Marshal(body)
	response.Body = string(content)

	return response
}

// InternalServerError 500
func (h HTTPHandler) InternalServerError() events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	var body responseBody
	body.Message = "Internal server error"
	response.StatusCode = 500

	content, _ := json.Marshal(body)
	response.Body = string(content)

	return response

}

// SuccessfulRequest 200
func (h HTTPHandler) SuccessfulRequest(body interface{}) events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	response.StatusCode = 200

	content, _ := json.Marshal(body)
	response.Body = string(content)

	return response
}

// BadRequest 400
func (h HTTPHandler) BadRequest(message string) events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	var body responseBody
	body.Message = message
	response.StatusCode = 400

	content, _ := json.Marshal(body)
	response.Body = string(content)

	return response
}

// NotFound 404
func (h HTTPHandler) NotFound(message string) events.APIGatewayProxyResponse {
	var response events.APIGatewayProxyResponse
	var body responseBody
	body.Message = message
	response.StatusCode = 404

	content, _ := json.Marshal(body)
	response.Body = string(content)

	return response
}
