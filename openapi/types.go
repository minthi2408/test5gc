package openapi

import (
	"net/url"
)

type Request struct {
	Path         string
	Method       string
	Body         interface{}
	HeaderParams map[string]string
	QueryParams  url.Values
	FormParams   url.Values
	FormFileName string
	FileName     string
	FileBytes    []byte
}

type Response struct {
	Raw  interface{}
	Body []byte
	Code int
}
