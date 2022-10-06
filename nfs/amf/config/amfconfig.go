package config

import (
	fabric_config "etrib5gc/fabric/config"
	"etrib5gc/sbi/models"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type NetworkFeatureSupport5GS struct {
	Enable  bool  `yaml:"enable" valid:"type(bool)"`
	Length  uint8 `yaml:"length" valid:"type(uint8)"`
	ImsVoPS uint8 `yaml:"imsVoPS" valid:"type(uint8)"`
	Emc     uint8 `yaml:"emc" valid:"type(uint8)"`
	Emf     uint8 `yaml:"emf" valid:"type(uint8)"`
	IwkN26  uint8 `yaml:"iwkN26" valid:"type(uint8)"`
	Mpsi    uint8 `yaml:"mpsi" valid:"type(uint8)"`
	EmcN3   uint8 `yaml:"emcN3" valid:"type(uint8)"`
	Mcsi    uint8 `yaml:"mcsi" valid:"type(uint8)"`
}

type Security struct {
	IntegrityOrder []string `yaml:"integrityOrder,omitempty" valid:"-"`
	CipheringOrder []string `yaml:"cipheringOrder,omitempty" valid:"-"`
}

type PlmnSupportItem struct {
	PlmnId     *models.PlmnId  `yaml:"plmnId" valid:"required"`
	SNssaiList []models.Snssai `yaml:"snssaiList,omitempty" valid:"required"`
}

type AmfConfig struct {
	Agent *fabric_config.AgentConfig

	AmfName    string
	NgapIpList []string //should be list of IP address

	//should go to context?
	GuamiList       []models.Guami
	TaiList         []models.Tai
	PlmnSupportList []PlmnItem
	DnnList         []string
	Locality        string

	NetworkName NetworkName
	Security    *Security

	T3502Value                      int
	T3512Value                      int
	Non3gppDeregistrationTimerValue int
	T3513                           TimerValue
	T3522                           TimerValue
	T3550                           TimerValue
	T3560                           TimerValue
	T3565                           TimerValue
	T3570                           TimerValue

	NetworkFeatureSupport5GS *NetworkFeatureSupport5GS
}

type PlmnItem struct {
	PlmnId     *models.PlmnId
	SNssaiList []models.Snssai
}

type NetworkName struct {
	Full  string `yaml:"full" valid:"type(string)"`
	Short string `yaml:"short,omitempty" valid:"type(string)"`
}

type TimerValue struct {
	Enable        bool          `yaml:"enable" valid:"type(bool)"`
	ExpireTime    time.Duration `yaml:"expireTime" valid:"type(time.Duration)"`
	MaxRetryTimes int           `yaml:"maxRetryTimes,omitempty" valid:"type(int)"`
}

func LoadConfig(f string) (*AmfConfig, error) {
	if _, err := ioutil.ReadFile(f); err != nil {
		return nil, err
	} else {

		var amfconf AmfConfig
		amfconf.AmfName = "amf"
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
		return &amfconf, nil
	}
}

func (c *AmfConfig) Get5gsNwFeatSuppEnable() bool {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.Enable
	}
	return true
}

func (c *AmfConfig) Get5gsNwFeatSuppImsVoPS() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.ImsVoPS
	}
	return 0
}

func (c *AmfConfig) Get5gsNwFeatSuppEmc() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.Emc
	}
	return 0
}

func (c *AmfConfig) Get5gsNwFeatSuppEmf() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.Emf
	}
	return 0
}

func (c *AmfConfig) Get5gsNwFeatSuppIwkN26() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.IwkN26
	}
	return 0
}

func (c *AmfConfig) Get5gsNwFeatSuppMpsi() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.Mpsi
	}
	return 0
}

func (c *AmfConfig) Get5gsNwFeatSuppEmcN3() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.EmcN3
	}
	return 0
}

func (c *AmfConfig) Get5gsNwFeatSuppMcsi() uint8 {
	if c.NetworkFeatureSupport5GS != nil {
		return c.NetworkFeatureSupport5GS.Mcsi
	}
	return 0
}

func convertTACConfigToModels(intString string) (string, error) {
	if tmp, err := strconv.ParseUint(intString, 10, 32); err != nil {
		return "", err
	} else {
		return fmt.Sprintf("%06x", tmp), nil
	}
}
