package openapi

import (
    "etri5gc/openapi/models"
)

type Error struct {
    status  string
    prob    *models.ProblemDetails
}

func NewError(status string, prob *models.ProblemDetails) *Error {
    return &Error{
        status: status,
        prob:   prob,
    }
}

func (e *Error) Error() string {
    return e.status
}

func (e *Error) Problem() *models.ProblemDetails {
    return e.prob
}
