package openapi

type ApiError struct {
	code   int
	status string
	data   interface{}
}

func NewApiError(code int, status string, data interface{}) *ApiError {
	return &ApiError{
		code:   code,
		status: status,
		data:   data,
	}
}

func (e *ApiError) Error() string {
	return e.status
}

func (e *ApiError) Code() int {
	return e.code
}

func (e *ApiError) Data() interface{} {
	return e.data
}
