package utils

import (
	"encoding/json"
)

func JsonString2map(str string) (jmap map[string]interface{}) {
	if err := json.Unmarshal([]byte(str), &jmap); err == nil {

	}
	return
}
