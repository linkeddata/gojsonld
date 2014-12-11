package gojsonld

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJSONFromFile(path string) (interface{}, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data, err := ReadJSON(file)
	return data, err
}

func ReadJSON(data []byte) (interface{}, error) {
	var jsonData interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
