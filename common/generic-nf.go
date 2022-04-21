package common

import "fmt"

type interface GenericNF {
	NFRegistration() error
}


type interface Service {
	Start() error
	Terminate()
}
