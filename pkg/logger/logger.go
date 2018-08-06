package logger

import (
	"encoding/json"
	"log"
)

// Logger is...
type Logger struct {
	code     logCode
	severity logSeverity
}

// NewLogger is
func NewLogger() Logger {

	defaultLogCode := logCode{
		InternalError:        "Internal error",
		JSONUnmarshal:        "Cannot unmarshal json",
		EnvVar:               "Cannot find environment variable",
		AuthorizerClaims:     "Invalid authorizer claims",
		URLPathParam:         "Invalid url path parameter",
		Authorization:        "Cannot find valid authorization for user",
		DBCreateItem:         "Cannot create item",
		DBGetItem:            "Cannot get item",
		DBUpdateItem:         "Cannot update item",
		DBDeleteItem:         "Cannot delete item",
		DBNotFoundItem:       "Cannot find item",
		DBExistingItem:       "Database item already exists",
		DBAttributeValue:     "Cannot find item attribute value",
		DBAttributeUnmarshal: "Cannot unmarshal item attribute value",
	}

	defaultLogSeverity := logSeverity{
		Info:    "SEVERITY_INFO",
		Warning: "SEVERITY_WARNING",
		Error:   "SEVERITY_ERROR",
	}

	return Logger{
		code:     defaultLogCode,
		severity: defaultLogSeverity,
	}
}

// accountLogEntry is
type logEntry struct {

	// identification data
	Message *string `json:"message,omitempty"`

	Severity string `json:"severity"`
	Code     string `json:"code"`
}

// WriteLogEntry is ..
func (l Logger) WriteLogEntry(message *string, severity string, code string) {

	// setting log flags for logging line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var le logEntry
	le.Message = message
	le.Severity = severity
	le.Code = code

	accountLogEntryBytes, err := json.Marshal(le)
	if err != nil {
		log.Println("Internal logger error, cannot marshal log entry, severity: " + l.GetErrorLogSeverity())
	} else {
		log.Println(string(accountLogEntryBytes))
	}
}
