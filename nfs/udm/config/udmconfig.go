package config

import (
	fabric_config "etrib5gc/fabric/config"
	//	"etrib5gc/openapi/models"
	//	"fmt"
	"io/ioutil"
	//	"strconv"
	//	"time"
)

type UdmConfig struct {
	Agent   *fabric_config.AgentConfig
	UdmName string
}

func LoadConfig(f string) (*UdmConfig, error) {
	if _, err := ioutil.ReadFile(f); err != nil {
		return nil, err
	} else {

		var udmconf UdmConfig
		udmconf.UdmName = "udm"
		/*
			if err := yaml.Unmarshal(content, &amfconf); err != nil {
				return nil, err
			}
			if err := amfconf.setDefaults(); err != nil {
				return nil, err
			} else {
				return &amfconf, nil
			}
		*/
		return &udmconf, nil
	}
}
