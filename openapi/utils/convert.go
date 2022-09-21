package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// ParameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func Param2Str(obj interface{}, format string) string {
	var delim string

	switch format {
	case "pipes":
		delim = "|"
	case "ssv":
		delim = " "
	case "tsv":
		delim = "\t"
	case "csv":
		delim = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delim, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}
