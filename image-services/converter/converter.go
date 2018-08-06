package converter

// Service is...
type Service struct {
	dbaccessor  DBAccessor
	httphandler HTTPHandler
	logger      Logger
	validator   Validator
}

// NewService is...
func NewService(dbac DBAccessor, http HTTPHandler, logg Logger, vald Validator) Service {
	return Service{
		dbaccessor:  dbac,
		httphandler: http,
		logger:      logg,
		validator:   vald,
	}
}
