package common

// the type of a network function (amf, smf, etc.)
type NetworkFunctionType string

// context field identity: PLMN, NSI, DNN etc (defined by a network operator)
type ContextFieldId string

// values of a context field identity (defined by a network operator)
type ContextFieldValue interface{}
type ContextFieldValues []ContextFieldValue

type NfContext map[ContextFieldId]ContextFieldValue

// context dicrionaty that enumerate all available ContextFieldId values and
// corresponding ContextFieldValue values
type ContextDict NfContext

func (c *NfContext) IsSubset(other *NfContext) bool {
	return true
}

func (c *NfContext) IsSuperset(other *NfContext) bool {
	return other.IsSubset(c)
}

// the abstraction of querries that are sent by an upper layer caller to select
// a remote producer
// Note: we have not sure how the query should be implemented yet. Should it be a
// complex sql-like query, or should it be any other simplified form?

type NfQuery interface {
	GetNfType() NetworkFunctionType
	Map2Context() (NfContext, error)
}
