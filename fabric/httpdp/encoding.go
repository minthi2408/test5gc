package httpdp

import (
	"etri5gc/openapi"
)

// a singleton
var oneEncoding *encoding = &encoding{}

// get the singlon object
func Encoding() *encoding {
	return oneEncoding
}

// emplement encoding and decoding of openapi's request/response to/from http's
// request/response
type encoding struct {
}

// unmarshaling the request body bytes (given a http request embeded in the
// given openapi's request)
func (enc *encoding) DecodeRequest(req *openapi.Request) error {
	return nil
}

// unmarshaling the response body bytes (give a http response embeded in the
// given openapi's response)
func (enc *encoding) DecodeResponse(req *openapi.Response) error {
	return nil
}

// marshaling the request body into bytes arrays the prepare a http request
func (enc *encoding) EncodeRequest(req *openapi.Request) error {
	return nil
}

// marshaling the response body into an byte array. No need to create an http
// response since gin.Context will write the response
func (enc *encoding) EncodeResponse(req *openapi.Response) error {
	return nil
}
