package shared

import (
	"fmt"
)

type ServiceError struct {
	Response             interface{}
	StatusCode           int
	ErrorCode            ErrorCode
	UseResponseAsMessage *bool
}

func (r *ServiceError) Error() string {
	useStandardError := r.UseResponseAsMessage != nil && *r.UseResponseAsMessage == false
	if r.UseResponseAsMessage == nil || useStandardError {
		return fmt.Sprintf("status_code=%d, error code=%s, error=%v", r.StatusCode, r.ErrorCode, r.Response)
	}

	return fmt.Sprintf("%v", r.Response)
}
