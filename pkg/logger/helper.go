package logger

type logCode struct {
	InternalError        string
	JSONUnmarshal        string
	EnvVar               string
	AuthorizerClaims     string
	URLPathParam         string
	Authorization        string
	DBCreateItem         string
	DBGetItem            string
	DBUpdateItem         string
	DBDeleteItem         string
	DBNotFoundItem       string
	DBExistingItem       string
	DBAttributeValue     string
	DBAttributeUnmarshal string
}

// GetInternalErrorLogCode is ...
func (l Logger) GetInternalErrorLogCode() string {
	return l.code.InternalError
}

// GetJSONUnmarshalLogCode is ...
func (l Logger) GetJSONUnmarshalLogCode() string {
	return l.code.JSONUnmarshal
}

// GetEnvVarLogCode is ...
func (l Logger) GetEnvVarLogCode() string {
	return l.code.EnvVar
}

// GetAuthorizerClaimsLogCode is ...
func (l Logger) GetAuthorizerClaimsLogCode() string {
	return l.code.AuthorizerClaims
}

// GetURLPathParamLogCode is ...
func (l Logger) GetURLPathParamLogCode() string {
	return l.code.URLPathParam
}

// GetAuthorizationLogCode is ...
func (l Logger) GetAuthorizationLogCode() string {
	return l.code.Authorization
}

// GetDBCreateItemLogCode is ...
func (l Logger) GetDBCreateItemLogCode() string {
	return l.code.DBCreateItem
}

// GetDBGetItemLogCode is ...
func (l Logger) GetDBGetItemLogCode() string {
	return l.code.DBGetItem
}

// GetDBUpdateItemLogCode is ...
func (l Logger) GetDBUpdateItemLogCode() string {
	return l.code.DBUpdateItem
}

// GetDBDeleteItemLogCode is ...
func (l Logger) GetDBDeleteItemLogCode() string {
	return l.code.DBDeleteItem
}

// GetDBNotFoundItemLogCode is ...
func (l Logger) GetDBNotFoundItemLogCode() string {
	return l.code.DBNotFoundItem
}

// GetDBExistingItemLogCode is ...
func (l Logger) GetDBExistingItemLogCode() string {
	return l.code.DBExistingItem
}

// GetDBAttributeValueLogCode is ...
func (l Logger) GetDBAttributeValueLogCode() string {
	return l.code.DBAttributeValue
}

// GetDBAttributeUnmarshalLogCode is ...
func (l Logger) GetDBAttributeUnmarshalLogCode() string {
	return l.code.DBAttributeUnmarshal
}

type logSeverity struct {
	Info    string
	Warning string
	Error   string
}

// GetInfoLogSeverity is ...
func (l Logger) GetInfoLogSeverity() string {
	return l.severity.Info
}

// GetWarningLogSeverity is ...
func (l Logger) GetWarningLogSeverity() string {
	return l.severity.Warning
}

// GetErrorLogSeverity is ...
func (l Logger) GetErrorLogSeverity() string {
	return l.severity.Error
}
