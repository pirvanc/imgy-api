package httphandler

type errorCode struct {
	DBItemNotFound      string
	DBCannotGetItem     string
	DBCannotUpdateItem  string
	DBCannotCreateItem  string
	DBCannotDeleteItem  string
	CannotDecodePayload string
	MissingPathParam    string
	DeletedItem         string
}

// GetDBItemNotFoundErrorCode is ...
func (h HTTPHandler) GetDBItemNotFoundErrorCode() string {
	return h.errorCode.DBItemNotFound
}

// GetDBCannotGetItemErrorCode is ...
func (h HTTPHandler) GetDBCannotGetItemErrorCode() string {
	return h.errorCode.DBCannotGetItem
}

// GetDBCannotUpdateItemErrorCode is ...
func (h HTTPHandler) GetDBCannotUpdateItemErrorCode() string {
	return h.errorCode.DBCannotUpdateItem
}

// GetDBCannotCreateItemErrorCode is ...
func (h HTTPHandler) GetDBCannotCreateItemErrorCode() string {
	return h.errorCode.DBCannotCreateItem
}

// GetDBCannotDeleteItemErrorCode is ...
func (h HTTPHandler) GetDBCannotDeleteItemErrorCode() string {
	return h.errorCode.DBCannotDeleteItem
}

// GetCannotDecodePayloadErrorCode is ...
func (h HTTPHandler) GetCannotDecodePayloadErrorCode() string {
	return h.errorCode.CannotDecodePayload
}

// GetMissingPathParamErrorCode is ...
func (h HTTPHandler) GetMissingPathParamErrorCode() string {
	return h.errorCode.MissingPathParam
}

// GetDeletedItemErrorCode is ...
func (h HTTPHandler) GetDeletedItemErrorCode() string {
	return h.errorCode.DeletedItem
}
