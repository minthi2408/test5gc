package utils

import (
	"fmt"
	"etri5gc/sbi/models"
)

func Param2String(v interface{}) string {
	//Dummy implementation
	return fmt.Sprintf("%v", v)
}

func String2ArrayOfstring(v string) []string {
	//dummy
	return []string{}
}

func String2UeContextInfoClass(s string) (v *models.UeContextInfoClass, err error) {
	//TODO
	return
}

func String2Guami(s string) (v *models.Guami, err error) {
	//TODO
	return
}
