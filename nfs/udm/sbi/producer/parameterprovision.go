package producer
import (
	"etri5gc/fabric/httpdp"
	"etri5gc/openapi"
	"etri5gc/openapi/models"
	udmprod "etri5gc/openapi/producers/udm"
	"fmt"
	"strings"
)

func parameterprovisionService(p *Producer) (service httpdp.HttpService) {
	fn := ginHandler
	ppdata := openapi.UDM_PP_DATA

	service.Routes = httpdp.HttpRoutes{
		{
			"Index",
			"GET",
			"/",
			httpdp.HttpIndexHandler,
		},
		{
			"HTTPUpdate",
			strings.ToUpper("Patch"),
			fmt.Sprintf("/:gpsi/%s", ppdata),
			fn(udmprod.Update, p),
		},
	}

	service.Group = openapi.UDM_PARAMETERPROVISION
	return
}

func (p *Producer) HandleUpdate(gpsi string, input *models.PpData) (prob *models.ProblemDetails) {
	return
}

