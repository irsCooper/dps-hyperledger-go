package utils

import "encoding/json"

func FromByteToStruct(data []byte, value interface{}) error {
	return json.Unmarshal(data, value)
}