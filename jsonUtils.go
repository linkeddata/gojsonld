package gojsonld

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSONFromFile(path string) (interface{}, error) {
	file, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		return nil, fileErr
	}
	var jsonData interface{}
	jsonErr := json.Unmarshal(file, &jsonData)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return jsonData, nil
}
