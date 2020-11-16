package util

import (
	"code.byted.org/gopkg/logs"
	"encoding/json"
)

func MarshallOrElseEmpty(v interface{}) string {
	if v == nil {
		return "nil"
	}
	data, err := json.Marshal(v)
	if err != nil {
		logs.Error("marshall exception, err: %v", err)
	}
	return string(data)
}
