package prodcontext

import (
	"etri5gc/openapi"
	"etri5gc/openapi/models"

	"github.com/gin-gonic/gin"
)

// an abstraction of the context where a request is received at a producer. The
// first handler method (openapi/producers) will inject a correct expecting
// body for decoding. The next handler (application layer) will process the
// decoded body then create a response. It then write it to a response writer
// to complete the whole procedure of handling a request.

type RequestContext interface {
	DecodeRequest(interface{}) //decode the request to get embeded body
	//	Problem() *models.ProblemDetails   //problem with decoding
	SetProblem(*models.ProblemDetails) //for the application handler to register a problem
	// SetResponse(*openapi.Response)     //set the response (after receiving it from the application handler
	//	WriteResponse()                    //write a prepared response to internal response writer
	Param(string) string // get a parameter from the request (application handler need it)
}

type AppProducerHandler func(RequestContext) *openapi.Response
type OpenApiProducerHandler func(RequestContext)

type ginContextEx struct {
	context  *gin.Context
	request  *openapi.Request
	response *openapi.Response
	problem  *models.ProblemDetails
}

func newGinContextEx(ctx *gin.Context) *ginContextEx {
	ret := &ginContextEx{
		context: ctx,
		request: &openapi.Request{
			Request: ctx.Request,
		},
	}
	return ret
}

//func (c *ginContextEx) Problem() *models.ProblemDetails {
//	return c.problem
//}

func (c *ginContextEx) SetProblem(prob *models.ProblemDetails) {
	c.problem = prob
}

//func (c *ginContextEx) SetResponse(resp *openapi.Response) {
//	c.response = resp
//}
func (c *ginContextEx) Param(key string) string {
	return c.context.Param(key)
}

//func (c *ginContextEx) Request() *openapi.Request {
//	return c.request
//}

func (c *ginContextEx) DecodeRequest(body interface{}) {
	if c.request == nil {
		return
	}
	//1. read the request body

	//2. decode the body (suppose that a right type of body has been pupulated
	c.request.Body = body
}

func (c *ginContextEx) WriteResponse() {
}

func BuildProducerRequestHandler(openapiFn OpenApiProducerHandler, appFn AppProducerHandler) gin.HandlerFunc {
	return func(context *gin.Context) {

		ctx := newGinContextEx(context)
		//call the openapi producer handler
		openapiFn(ctx)
		//check the problem details
		if ctx.problem == nil {
			appFn(ctx)
		}
		ctx.WriteResponse()
	}
}
