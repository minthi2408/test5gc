package config
import (
    "time"
	"etri5gc/openapi/models"
    fabric_config "etri5gc/fabric/config"
)
type AmfConfig struct {
    Agent       *fabric_config.AgentConfig

    AmfName     string
    NgapIpList  []string //should be list of IP address

    //should go to context?
    GuarmiList    []models.Guami
    TaiList    []models.Tai
    PlmnSupportList     []PlmnItem
    DnnList     []string
    Locality    string

    NetworkName     NetworkName
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

