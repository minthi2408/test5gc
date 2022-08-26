package prodcontext

import (
    "net/http"
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
	DecodeRequest(body interface{}) //decode the request to get embeded body
//	SetProblem(*models.ProblemDetails) //for the application handler to register a problem
	Param(string) string // get a parameter from the request (application handler need it)
}

type AppProducerHandler func(RequestContext) *openapi.Response
type OpenApiProducerHandler func(RequestContext)

type ginContextEx struct {
	context  *gin.Context
	request  *openapi.Request
	response *openapi.Response
//	problem  *models.ProblemDetails
    encoding   openapi.ProducerEncoding 
}

func newGinContextEx(ctx *gin.Context, encoding openapi.ProducerEncoding) *ginContextEx {
	ret := &ginContextEx{
		context: ctx,
		request: &openapi.Request{
			Request: ctx.Request,
		},
	}
	return ret
}
/*
func (c *ginContextEx) SetProblem(prob *models.ProblemDetails) {
	c.problem = prob
}
*/
func (c *ginContextEx) Param(key string) string {
	return c.context.Param(key)
}

//must be called by an openapi's producer handler
func (c *ginContextEx) DecodeRequest(body interface{}) {
	//1. read the request body
    var err error
    if c.request.BodyBytes, err = c.context.GetRawData(); err != nil {
        c.response = &openapi.Response {
            StatusCode: http.StatusInternalServerError,
            Body: &models.ProblemDetails{
            },
        }
    } else {
	    //2. decode the body (suppose that a right type of body has been pupulated
    	c.request.Body = body
        if err = c.encoding.DecodeRequest(c.request); err != nil {
            c.response = &openapi.Response {
                StatusCode: http.StatusBadRequest,
                Body: &models.ProblemDetails{
                },
            }
        }
    }
}

func (c *ginContextEx) writeResponse() {
    if err:=c.encoding.EncodeResponse(c.response); err != nil {
        c.context.JSON(http.StatusInternalServerError, &models.ProblemDetails {
        })
    } else {
        //TODO: how to set the Content-Type?
        c.context.Data(c.response.StatusCode, "application/json", c.response.BodyBytes)
    }
}

func BuildProducerRequestHandler(openapiFn OpenApiProducerHandler, appFn AppProducerHandler, encoding openapi.ProducerEncoding) gin.HandlerFunc {
	return func(context *gin.Context) {

		ctx := newGinContextEx(context, encoding)
		//call the openapi producer handler to set the request body to decode
		openapiFn(ctx)
		if ctx.response == nil { //no error found after decoding
			ctx.response = appFn(ctx) // call the application handler to handle the decoded request
		}
		ctx.writeResponse()
	}
}
