package fabric

type NFType string

// this class represent a query to find a network function
type NFContext interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
}
