package sbi

import (
	"fmt"
)

type Config struct {
	Scheme       string `yaml:"scheme" valid:"required,scheme"`
	RegisterIPv4 string `yaml:"registerIPv4,omitempty" valid:"required,host"` // IP that is registered at NRF.
	BindingIPv4  string `yaml:"bindingIPv4,omitempty" valid:"required,host"`  // IP used to run the server in the node.
	Port         int    `yaml:"port,omitempty" valid:"required,port"`
	Tls          *Tls   `yaml:"tls,omitempty" valid:"optional"`
}

func (c *Config) validate() (bool, error) {
	govalidator.TagMap["scheme"] = govalidator.Validator(func(str string) bool {
		return str == "https" || str == "http"
	})

	if tls := c.Tls; tls != nil {
		if result, err := tls.validate(); err != nil {
			return result, err
		}
	}

	if _, err := govalidator.ValidateStruct(c); err != nil {
		return false, AppendInvalid(err)
	}

	return true, nil
}

type Tls struct {
	Pem string `yaml:"pem,omitempty" valid:"type(string),minstringlength(1),required"`
	Key string `yaml:"key,omitempty" valid:"type(string),minstringlength(1),required"`
}

func (t *Tls) validate() (bool, error) {
	result, err := govalidator.ValidateStruct(t)
	return result, err
}

func AppendInvalid(err error) error {
	var errs govalidator.Errors

	es := err.(govalidator.Errors).Errors()
	for _, e := range es {
		errs = append(errs, fmt.Errorf("Invalid %w", e))
	}

	return error(errs)
}

