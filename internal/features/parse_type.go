package features

import (
	"strconv"
	"strings"
)

func ConvType(data string) interface{} {
	isString := strings.Contains(data, "\"")
	if isString {
		return data
	}

	digital, err := strconv.ParseInt(data, 10, 64)
	if err == nil {
		return digital
	}

	if data == "false" {
		return false
	} else if data == "true" {
		return true
	}

	return data
}
