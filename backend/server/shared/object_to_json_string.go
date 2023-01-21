package shared

import "encoding/json"

func ToJSONString(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
