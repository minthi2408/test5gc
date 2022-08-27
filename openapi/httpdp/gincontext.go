package httpdp

import (
    "net/http"
	fabricdp "etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"

	"github.com/gin-gonic/gin"
)

//a wrapper for gin.Context. It implements the openapi.RequestContext
//abstraction
type ginContextEx struct {
    context  *gin.Context 
    request  *openapi.Request
	response *openapi.Response
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
        if err = fabricdp.Encoding().DecodeRequest(c.request); err != nil {
            c.response = &openapi.Response {
                StatusCode: http.StatusBadRequest,
                Body: &models.ProblemDetails{
                },
            }
        }
    }
}

func (c *ginContextEx) writeResponse() {
    if err:=fabricdp.Encoding().EncodeResponse(c.response); err != nil {
        //write the poblem in application/json format
        c.context.JSON(http.StatusInternalServerError, &models.ProblemDetails {
        })
    } else {
        //TODO: how to set the Content-Type?
        c.context.Data(c.response.StatusCode, "application/json", c.response.BodyBytes)
    }
}

//build a service handler for gin
func BuildProducerRequestHandler(openapiFn openapi.OpenApiProducerHandler, appFn openapi.AppProducerHandler) gin.HandlerFunc {
	return func(context *gin.Context) {

		ctx := newGinContextEx(context)
		//call the openapi producer handler to set the request body to decode
		openapiFn(ctx)
		if ctx.response == nil { //no error found after decoding
			ctx.response = appFn(ctx) // call the application handler to handle the decoded request
		}
		ctx.writeResponse()
	}
}
