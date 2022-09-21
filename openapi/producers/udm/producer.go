package udm

import "etri5gc/openapi/models"

//NF should implement this interface
type UdmProducer interface {
	//event exposure
	HandleCreateEeSubscription(string, *models.EeSubscription) (models.CreatedEeSubscription, *models.ProblemDetails)
	HandleDeleteEeSubscription(string, string) *models.ProblemDetails
	HandleUpdateEeSubscription(string, string, []models.PatchItem) *models.ProblemDetails


	//parameter provision
	HandleUpdate(string, *models.PpData) *models.ProblemDetails
}
