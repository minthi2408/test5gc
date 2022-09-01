package openapi

import (
	"etri5gc/fabric/common"
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

	BodyBytes []byte      //encoded body
	Request   interface{} //dataplane protocol-bound request
}

func (msg *Request) MsgType() common.ServiceMsgType {
	return common.SERVICE_MSG_TYPE_OPENAPI
}

func DefaultRequest() *Request {
	ret := &Request{
		HeaderParams: make(map[string]string),
	}
	ret.HeaderParams["Content-Type"] = "application/json"
	ret.HeaderParams["Accept"] = "application/json;application/problem+json"
	return ret
}

type Response struct {
	Response  interface{}
	BodyBytes []byte

	Body       interface{}
	Status     string
	StatusCode int
}

func (msg *Response) MsgType() common.ServiceMsgType {
	return common.SERVICE_MSG_TYPE_OPENAPI
}

// Abstraction of a consumer client
type ConsumerClient interface {
	//Encode the request into a data-plane specific format then send
	Send(*Request) (*Response, error)
	//Populate the openapi response by decoding the received response (in data-plane
	//specific format).
	DecodeResponse(*Response) error
}

// Abtraction of encoding methods for a producer
type ProducerEncoding interface {
	DecodeRequest(*Request) error
	EncodeResponse(*Response) error
}

// an abstraction of the context where a request is received at a producer. The
// first handler method (openapi/producers) will inject a correct expected
// body for decoding. The next handler (application layer) will process the
// decoded body then create a response. It then write it to a response writer
// to complete the whole procedure of handling a request.

type RequestContext interface {
	DecodeRequest(body interface{}) //decode the request to get embeded body
	Param(string) string            // get a parameter from the request (application handler need it)
}

//abstraction for application handlers where signaling procedures are
//implemented
type AppProducerHandler func(RequestContext) *Response

// abstraction for openapi producer handlers where correct expected
// data structures should be populated for being ready to decode from a received response
type OpenApiProducerHandler func(RequestContext)
