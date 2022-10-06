package config

import (
	fabric_config "etri5gc/fabric/config"
	//	"etri5gc/openapi/models"
	//	"fmt"
	"io/ioutil"
	//	"strconv"
	//	"time"
)

type PcfConfig struct {
	Agent   *fabric_config.AgentConfig
	PcfName string
}

func LoadConfig(f string) (*PcfConfig, error) {
	if _, err := ioutil.ReadFile(f); err != nil {
		return nil, err
	} else {

		var pcfconf PcfConfig
		pcfconf.PcfName = "pcf"
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
		return &pcfconf, nil
	}
}
