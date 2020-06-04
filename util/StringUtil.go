package util

import "encoding/json"

func StrToJson(jsonStr string) map[string]interface{} {

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dat); err == nil {

	}
	return dat
}
