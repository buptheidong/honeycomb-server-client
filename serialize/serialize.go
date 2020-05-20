package serialize

import (
	"encoding/json"
)

//JSONSerialize Json序列化
func JSONSerialize(value interface{}) string {
	str, err := json.Marshal(value)
	if err != nil {
		panic(err.Error())
	}
	return string(str)
}

//JSONDeserialize Json反序列化
func JSONDeserialize(result interface{}, jsonStr string) interface{} {
	err := json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		panic(err.Error())
	}
	return result
}