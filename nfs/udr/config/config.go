package config

import (
	fabric_config "etri5gc/fabric/config"
	//	"etri5gc/openapi/models"
	//	"fmt"
	"io/ioutil"
	//	"strconv"
	//	"time"
)

type UdrConfig struct {
	Agent   *fabric_config.AgentConfig
	UdrName string
}

func LoadConfig(f string) (*UdrConfig, error) {
	if _, err := ioutil.ReadFile(f); err != nil {
		return nil, err
	} else {

		var udrconf UdrConfig
		udrconf.UdrName = "udr"
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
		return &udrconf, nil
	}
}
