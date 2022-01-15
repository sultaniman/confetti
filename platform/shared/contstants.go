package shared

type ErrorCode string

const (
	// HTTP errors

	InvalidPayload ErrorCode = "invalid_payload"
	NotFound       ErrorCode = "not_found"
	Conflict       ErrorCode = "conflict"
	Unauthorized   ErrorCode = "unauthorized"
	Forbidden      ErrorCode = "forbidden"
	BadRequest     ErrorCode = "bad_request"
	ServerError    ErrorCode = "internal_error"
	UpdateError    ErrorCode = "update_error"
	InactiveUser   ErrorCode = "inactive_user"
)