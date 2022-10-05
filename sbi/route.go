package sbi

type SbiRoute struct {
	Label   string
	Method  string
	Path    string
	Handler func(RequestContext, interface{}) Response
}

type SbiRoutes []SbiRoute
