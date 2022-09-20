package udm
import (
	"etri5gc/nfs/udm/config"
)

type UdmContext struct {
	conf	*config.UdmConfig
	uelist	UdmUeList
}




func New(conf *config.UdmConfig) *UdmContext {
	ret := &UdmContext{
		conf:	conf,
	}
	return ret
}
