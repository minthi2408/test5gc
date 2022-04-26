package sbi

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

func AddHttpRoutes(engine *gin.Engine, groupname string, routes Routes) *gin.RouterGroup{
	group := engine.Group(groupname)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// IndexHandler is the index handler.
func IndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Etri5GC!")
}
