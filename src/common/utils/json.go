package utils

import "encoding/json"

func Marshal(i interface{}) string {
	res, _ := json.Marshal(i)
	return string(res)
}
