package fabric

//an abstract profile of a network function
type NFProfile interface {
	Identity() NFId
}

//identity of a network function
type NFId string

//a load balancer interface
type Selector interface {
	Select([]NFProfile) NFId
}
