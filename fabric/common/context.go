package common

//the type of a network function (amf, smf, etc.)
type NetworkFunctionType string

// the abstraction of querries that are sent by an upper layer caller to select
// a remote producer
// Note: we have not sure how the query should be implemented yet. Should it be a
// complex sql-like query, or should it be any other simplified form?

type NfQuery interface {
	GetNfType() NetworkFunctionType
	Map2Context() (NfContext, error)
}

type NfContextKey string

type NfContextValue struct {
	//value type
	//validator
}

type NfContextField struct {
	id    NfContextKey
	value []NfContextValue
}

//define a context where a network function operate
//a NfQuery should be map to a single NfContext so that a list of currently
//deployed (running) NFs can be retrieved

type NfContext struct {
	fields []NfContextField
}

func (c *NfContext) IsSubset(other *NfContext, dict *NfContextDict) bool {
	return true
}
func (c *NfContext) IsSuperset(other *NfContext, dict *NfContextDict) bool {
	return false
}

type NfContextDict NfContext

/*
type NfContext interface {
	IsSubset(NfContext, *NfContextDict) bool
	IsSuperset(NfContext, *NfContextDict) bool
}
*/
