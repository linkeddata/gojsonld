package gojsonld

import (
	"reflect"
	"sort"
	"strings"
)

const MAX_CONTEXT_URLS = 10

var (
	allKeywords = map[string]bool{
		"@base":        true,
		"@context":     true,
		"@container":   true,
		"@default":     true,
		"@embed":       true,
		"@explicit":    true,
		"@graph":       true,
		"@id":          true,
		"@index":       true,
		"@language":    true,
		"@list":        true,
		"@omitDefault": true,
		"@reverse":     true,
		"@preserve":    true,
		"@set":         true,
		"@type":        true,
		"@value":       true,
		"@vocab":       true,
	}
)

func isKeyword(key interface{}) bool {
	switch s := key.(type) {
	case string:
		return allKeywords[s]
	}
	return false
}

func isScalar(value interface{}) bool {
	_, isString := value.(string)
	_, isFloat64 := value.(float64)
	_, isFloat32 := value.(float32)
	_, isInt64 := value.(int64)
	_, isInt32 := value.(int32)
	_, isBoolean := value.(bool)
	if isString || isFloat32 || isFloat64 || isInt32 ||
		isInt64 || isBoolean {
		return true
	}
	return false
}

func isValueObject(value interface{}) bool {
	valueMap, isMap := value.(map[string]interface{})
	_, containsValue := valueMap["@value"]
	if isMap && containsValue {
		return true
	}
	return false
}

func isValidValueObject(value interface{}) bool {
	valueMap, isMap := value.(map[string]interface{})
	if !isMap {
		return false
	}
	if len(valueMap) > 4 {
		return false
	}
	for key := range valueMap {
		if key != "@value" && key != "@language" &&
			key != "@type" && key != "@index" {
			return false
		}
	}
	_, hasLanguage := valueMap["@language"]
	_, hasType := valueMap["@type"]
	if hasLanguage && hasType {
		return false
	}
	return true
}

func isListObject(value interface{}) bool {
	valueMap, isMap := value.(map[string]interface{})
	_, containsList := valueMap["@list"]
	if isMap && containsList {
		return true
	}
	return false
}

func isNil(value interface{}) bool {
	switch value.(type) {
	case string, int64, int32, float64, float32, bool:
		return false
	}

	if value == nil || reflect.ValueOf(value).IsNil() {
		return true
	} else {
		return false
	}
}

func deepCompareMatters(v1, v2 interface{}, listOrderMatters bool) bool {
	return reflect.DeepEqual(v1, v2)
}

func deepCompare(v1, v2 interface{}) bool {
	return deepCompareMatters(v1, v2, false)
}

func deepContains(values []interface{}, value interface{}) bool {
	for _, item := range values {
		if deepCompare(item, value) {
			return true
		}
	}
	return false
}

func deepCopy(value interface{}) interface{} {
	switch v := value.(type) {
	case string, int64, int32, float64, float32:
		valueCopy := v
		return valueCopy
	case []interface{}:
		tmpArray := make([]interface{}, 0)
		for _, item := range v {
			tmpArray = append(tmpArray, deepCopy(item))
		}
		return tmpArray
	case map[string]interface{}:
		tmpMap := make(map[string]interface{}, 0)
		for key, item := range v {
			tmpMap[key] = deepCopy(item)
		}
		return tmpMap
	}
	return nil
}

func mergeValue(obj map[string]interface{}, key string, value interface{}) {
	if obj == nil {
		return
	}

	values, ex := obj[key].([]interface{})
	if !ex {
		values = make([]interface{}, 0)
	}

	if key == "@list" || isListObject(value) ||
		!deepContains(values, value) {
		values = append(values, value)
		obj[key] = values
		return
	}
}

type InverseSlice []string

func (is InverseSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is InverseSlice) Len() int {
	return len(is)
}

func compareShortestLeast(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return len(s1) < len(s2)
	} else {
		return s1 < s2
	}
}

func (is InverseSlice) Less(i, j int) bool {
	s1, s2 := is[i], is[j]
	return compareShortestLeast(s1, s2)
}

func specialSortInverse(keys []string) {
	sort.Sort(InverseSlice(keys))
}

func sortedKeys(inputMap map[string]interface{}) []string {
	keys := make([]string, 0)
	for key := range inputMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func isAbsoluteIri(value string) bool {
	// TODO: this is a bit simplistic!
	return strings.Contains(value, ":")
}

func isIRI(value interface{}) bool {
	//TODO improve function
	valueString, isString := value.(string)
	if !isString {
		return false
	}
	if strings.HasPrefix(valueString, "_") {
		return false
	}
	if !strings.Contains(valueString, ":") {
		return false
	}
	return true
}

func isRelativeIri(value string) bool {
	if !(isKeyword(value) || isAbsoluteIri(value)) {
		return true
	}
	return false
}

func isNodeObject(value interface{}) bool {
	valueMap, isMap := value.(map[string]interface{})
	if !isMap {
		return false
	}
	_, hasValue := valueMap["@value"]
	_, hasList := valueMap["@list"]
	_, hasSet := valueMap["@set"]
	if !(hasValue || hasList || hasSet) {
		return true
	}
	return false
}

func isBlankNodeIdentifier(value string) bool {
	if strings.HasPrefix(value, "_:") {
		return true
	}
	return false
}

func convertFloatValue(value string) string {
	minusIndex := strings.Index(value, "-")
	plusIndex := strings.Index(value, "+")
	var index int
	if minusIndex > plusIndex {
		index = minusIndex
	} else {
		index = plusIndex
	}
	base := value[:(index - 1)]
	dotIndex := strings.Index(base, ".")
	if dotIndex < 0 {
		base += ".0"
	}
	exponent := value[(index + 1):]
	exponent = strings.TrimLeft(exponent, "0")
	if exponent == "" {
		exponent = "0"
	}
	if plusIndex > -1 {
		return base + "E" + exponent
	} else {
		return base + "E-" + exponent
	}
}
