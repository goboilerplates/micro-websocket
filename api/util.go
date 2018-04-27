package api

import "encoding/json"

// ConvertInterfaceToJSONBytes converts any interface to json bytes.
func ConvertInterfaceToJSONBytes(input interface{}, inputError error) []byte {
	if inputError != nil {
		return nil
	}
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil
	}
	return bytes
}

// ConvertBytesToMap converts byte array into map interface.
func ConvertBytesToMap(input []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(input, &data)
	return data, err
}
