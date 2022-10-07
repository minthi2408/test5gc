package consumer

import (
	"etrib5gc/sbi"
)

type UdmConsumer struct {
	cli sbi.ConsumerClient //connection to a remote UDM
	//other attributes to maintain remote context
}

func NewUdmConsumer(cli sbi.ConsumerClient) *UdmConsumer {
	return &UdmConsumer{
		cli: cli,
	}
}

//example high-level consumer api
func (c *UdmConsumer) SendAbcd(param1 string, param2 int32) error {
	//call the low level Sbi consumer API
	//udmee.{{methodname}}(cli, {{method-params}})
	return nil
}
