package strconv

import "encoding/json"

// StringToJSONArray convert a json string to a string array
func StringToJSONArray(value string) []string {
	var data []string
	_ = json.Unmarshal([]byte(value), &data)
	return data
}
