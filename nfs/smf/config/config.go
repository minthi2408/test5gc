package config

import (
	fabric_config "etri5gc/fabric/config"
	//	"etri5gc/openapi/models"
	//	"fmt"
	"io/ioutil"
	//	"strconv"
	//	"time"
)

type SmfConfig struct {
	Agent   *fabric_config.AgentConfig
	SmfName string
}

func LoadConfig(f string) (*SmfConfig, error) {
	if _, err := ioutil.ReadFile(f); err != nil {
		return nil, err
	} else {

		var smfconf SmfConfig
		smfconf.SmfName = "smf"
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
		return &smfconf, nil
	}
}
