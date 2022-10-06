package config

import (
	fabric_config "etrib5gc/fabric/config"
	"io/ioutil"
)

type AusfConfig struct {
	Agent    *fabric_config.AgentConfig
	AusfName string
}

func LoadConfig(f string) (*AusfConfig, error) {
	if _, err := ioutil.ReadFile(f); err != nil {
		return nil, err
	} else {

		var ausfconf AusfConfig
		ausfconf.AusfName = "ausf"
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
		return &ausfconf, nil
	}
}
