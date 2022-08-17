package fabric

//represent a service agent for an NF
type ServiceAgent interface {
	Start()
}

/*
type Registration interface {
	Beat() error //send an hearbeat
	Register() error //register the underlying service
}


type Discovery interface {
	Search(ServiceQuery) (ServiceProfiles, error)
}


//definition of a service query
type ServiceQuery map[string]string

//load balancer select one of the service instances
type LoadBalancer interface {
	Select(ServiceProfiles) ServiceProfile
}
*/
