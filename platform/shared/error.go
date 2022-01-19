package shared

type ErrorDetail struct {
	Field string
	Value string
}

type HTTPError struct {
	Code    ErrorCode
	Message string
	Details interface{} `swaggertype:"object"`
}

func ServerErrorResponse(err error, code ErrorCode) HTTPError {
	responseCode := ServerError
	if code == "" {
		responseCode = code
	}

	return HTTPError{
		Code:    responseCode,
		Message: err.Error(),
	}
}
