package sbi

import (
	"etrib5gc/sbi/models"
)

type ApiError struct {
	code   int
	status string
	data   interface{}
}

func NewApiError(code int, status string, data interface{}) *ApiError {
	return &ApiError{
		code:   code,
		status: status,
		data:   data, //either a ProblemDetails or an json-marshallable api-specific error
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

func ApiErrFromProb(prob *models.ProblemDetails) *ApiError {
	return NewApiError(int(prob.Status), prob.Detail, prob)
}

func (e *ApiError) Problem() (prob *models.ProblemDetails, ok bool) {
	prob, ok = e.data.(*models.ProblemDetails)
	return
}
