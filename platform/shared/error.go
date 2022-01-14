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

func InvalidPayloadResponse(err error, code ErrorCode) HTTPError {
	responseCode := InvalidPayload
	if code == "" {
		responseCode = code
	}
	return HTTPError{
		Code:    responseCode,
		Message: err.Error(),
	}
}

func NotFoundResponse(msg string, code ErrorCode) HTTPError {
	responseCode := NotFound
	if code == "" {
		responseCode = code
	}

	return HTTPError{
		Code:    responseCode,
		Message: msg,
	}
}

func UnauthorizedResponse(msg string, code ErrorCode) HTTPError {
	responseCode := Unauthorized
	if code == "" {
		responseCode = code
	}

	return HTTPError{
		Code:    responseCode,
		Message: msg,
	}
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
