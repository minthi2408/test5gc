package httpdp

import (
    "etri5gc/openapi"
)

var oneEncoding *encoding = &encoding{}

func Encoding() *encoding {
    return oneEncoding
}

type encoding struct{
}
//unmarshaling the request body bytes (given a http request)
func (enc *encoding) DecodeRequest(req *openapi.Request) error {
    return nil
}
//unmarshaling the response body bytes (give a http response)
func (enc *encoding) DecodeResponse(req *openapi.Response) error {
    return nil
}

//marshaling the request body into bytes arrays the prepare a http request
func (enc *encoding) EncodeRequest(req *openapi.Request) error {
    return nil
}

//marshaling the response body into an byte array. No need to create an http
//response since gin.Context will write the response 
func (enc *encoding) EncodeResponse(req *openapi.Response) error {
    return nil
}
