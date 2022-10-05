package models

type NfProfile struct{}
type SubscribedData struct{}
type SubscribedSnssai struct{}

type PostSmContextsResponse struct {
	JsonData SmContextCreatedData `json:"jsonData,omitempty"`

	BinaryDataN2SmInformation []byte `json:"binaryDataN2SmInformation,omitempty"`
}

type UpdateSmContextResponse struct {
	JsonData SmContextUpdatedData `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage []byte `json:"binaryDataN1SmMessage,omitempty"`

	BinaryDataN2SmInformation []byte `json:"binaryDataN2SmInformation,omitempty"`
}
type UpdateSmContextErrorResponse struct {
	JsonData SmContextUpdateError `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage []byte `json:"binaryDataN1SmMessage,omitempty"`

	BinaryDataN2SmInformation []byte `json:"binaryDataN2SmInformation,omitempty"`
}
type PostSmContextsErrorResponse struct {
	JsonData SmContextCreateError `json:"jsonData,omitempty"`

	BinaryDataN1SmMessage []byte `json:"binaryDataN1SmMessage,omitempty"`

	BinaryDataN2SmMessage []byte `json:"binaryDataN2SmMessage,omitempty"`
}
type UEContextTransferResponse struct {
	JsonData UeContextTransferRspData `json:"jsonData,omitempty"`

	BinaryDataN2Information []byte `json:"binaryDataN2Information,omitempty"`

	BinaryDataN2InformationExt1 []byte `json:"binaryDataN2InformationExt1,omitempty"`
}
