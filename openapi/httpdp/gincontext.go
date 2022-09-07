package httpdp

import (
	fabricdp "etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// a wrapper for gin.Context. It implements the openapi.RequestContext
// abstraction
type ginContextEx struct {
	context  *gin.Context
	request  *openapi.Request
	response openapi.Response
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

func (c *ginContextEx) RequestBody() interface{} {
	return c.request.Body
}

// must be called by an openapi's producer handler
func (c *ginContextEx) DecodeRequest(body interface{}) *models.ProblemDetails {
	//1. read the request body
	var err error
	if c.request.BodyBytes, err = c.context.GetRawData(); err != nil {
		return &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		}
	} else {
		//2. decode the body (suppose that a right type of body has been pupulated
		c.request.Body = body
		if err = fabricdp.Encoding().DecodeRequest(c.request); err != nil {
			return &models.ProblemDetails{
				Status: http.StatusBadRequest,
				Detail: err.Error(),
			}
		}
	}
	return nil
}

func (c *ginContextEx) writeResponse() {
	//encode the body of the response into bytes
	if err := fabricdp.Encoding().EncodeResponse(&c.response); err != nil {
		//write the poblem in application/json format
		c.context.JSON(http.StatusInternalServerError, &models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
		})
	} else {
		//TODO: how to set the Content-Type?
		//write to the response writer
		c.context.Data(c.response.StatusCode, "application/json", c.response.BodyBytes)
	}
}

// build a service handler for gin
// the first parameter is the openapi handler, the second parameter is an
// NF specific producer implementation.
func CreateGinHandler(openapiFn openapi.OpenApiProducerHandler, handler interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {

		ctx := newGinContextEx(context)
		//call the openapi producer handler to decode request and call
		//application handler to process the request
		ctx.response = openapiFn(ctx, handler)
		//write the response or error to the dataplane
		ctx.writeResponse()
	}
}
