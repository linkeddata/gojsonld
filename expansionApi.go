package gojsonld

import (
	"strings"
)

func expand(activeContext *Context, activeProperty *string,
	element interface{}) (interface{}, error) {
	//1)
	if isNil(element) {
		return nil, nil
	}
	// 2)
	if isScalar(element) {
		if isNil(activeProperty) || *activeProperty == "@graph" {
			return nil, nil
		}
		expandedValue, expandErr := expandValue(activeContext, *activeProperty, element)
		return expandedValue, expandErr
	}
	// 3)
	if elementArray, isArray := element.([]interface{}); isArray {
		// 3.1)
		result := make([]interface{}, 0)
		for _, item := range elementArray {
			// 3.2.1)
			expandedItem, expandErr := expand(activeContext, activeProperty, item)
			if !isNil(expandErr) {
				return nil, expandErr
			}
			// 3.2.2)
			expandedArray, isArray := expandedItem.([]interface{})
			if !isNil(activeProperty) && ((*activeProperty == "@list" ||
				activeContext.getContainer(*activeProperty) == "@list") &&
				(isArray || isListObject(expandedItem))) {
				return nil, LIST_OF_LISTS
			}
			// 3.2.3)
			if isArray {
				for _, item := range expandedArray {
					result = append(result, item)
				}
			} else if !isNil(expandedItem) {
				result = append(result, expandedItem)
			}
		}
		// 3.3)
		return result, nil
	}
	// 4)
	elementMap := element.(map[string]interface{})
	// 5)
	if context, hasContext := elementMap["@context"]; hasContext {
		emptyArray := make([]string, 0)
		processedContext, processErr := parse(activeContext, context, emptyArray)
		if !isNil(processErr) {
			return nil, processErr
		}
		activeContext = processedContext
	}
	// 6)
	result := make(map[string]interface{}, 0)
	//7
	keys := sortedKeys(elementMap)
	for _, key := range keys {
		value := elementMap[key]
		// 7.1)
		if key == "@context" {
			continue
		}
		// 7.2)
		expandedProperty, expandErr := expandIri(activeContext, &key,
			false, true, nil, nil)
		if !isNil(expandErr) {
			return nil, expandErr
		}
		var expandedValue interface{} = nil
		// 7.3)
		if isNil(expandedProperty) || (!strings.Contains(*expandedProperty, ":") &&
			!isKeyword(*expandedProperty)) {
			continue
		}
		//7.4)
		if isKeyword(*expandedProperty) {
			// 7.4.1)
			if !isNil(activeProperty) && *activeProperty == "@reverse" {
				return nil, INVALID_REVERSE_PROPERTY_MAP
			}
			// 7.4.2)
			if _, hasProperty := result[*expandedProperty]; hasProperty {
				return nil, COLLIDING_KEYWORDS
			}
			// 7.4.3)
			if valueString, isString := value.(string); !isString &&
				*expandedProperty == "@id" {
				return nil, INVALID_ID_VALUE
			} else {
				tmpExpandedValue, expandedValueErr := expandIri(activeContext,
					&valueString, true, false, nil, nil)
				if !isNil(expandedValueErr) {
					return nil, expandedValueErr
				}
				expandedValue = *tmpExpandedValue
			}
			// 7.4.4)
			valueString, isString := value.(string)
			valueArray, isArray := value.([]interface{})
			valueMap, isMap := value.(map[string]interface{})
			if *expandedProperty == "@type" {
				if isString {
					tmpExpandedValue, expandErr := expandIri(activeContext,
						&valueString, true, true, nil, nil)
					if !isNil(expandErr) {
						return nil, expandErr
					}
					expandedValue = *tmpExpandedValue
				} else if isArray {
					expandedArray := make([]interface{}, 0)
					for _, item := range valueArray {
						itemString := item.(string)
						tmpExpandedValue, expandErr := expandIri(activeContext,
							&itemString, true, true, nil, nil)
						if !isNil(expandErr) {
							return nil, expandErr
						}
						expandedArray = append(expandedArray, *tmpExpandedValue)
					}
					expandedValue = expandedArray
				} else if isMap {
					//TODO check if empty map check should be part of the spec
					if len(valueMap) != 0 {
						return nil, INVALID_TYPE_VALUE
					}
					expandedValue = value
				} else {
					return nil, INVALID_TYPE_VALUE
				}
			}
			// 7.4.5)
			if *expandedProperty == "@graph" {
				graphArg := "@graph"
				tmpExpandedValue, expandErr := expand(activeContext,
					&graphArg, value)
				if !isNil(expandErr) {
					return nil, expandErr
				}
				expandedValue = tmpExpandedValue
			}
			// 7.4.6)
			if *expandedProperty == "@value" {
				if !(isNil(value) || isScalar(value)) {
					return nil, INVALID_VALUE_OBJECT_VALUE
				}
				expandedValue = value
				if isNil(expandedValue) {
					result["@value"] = nil
					continue
				}
			}
			// 7.4.7)
			if *expandedProperty == "@language" {
				if !isString {
					return nil, INVALID_LANGUAGE_TAGGED_STRING
				}
				expandedValue = strings.ToLower(valueString)
			}
			// 7.4.8)
			if *expandedProperty == "@index" {
				if !isString {
					return nil, INVALID_INDEX_VALUE
				}
				expandedValue = value
			}
			// 7.4.9)
			if *expandedProperty == "@list" {
				// 7.4.9.1)
				if isNil(activeProperty) || *activeProperty == "@graph" {
					continue
				}
				// 7.4.9.2)
				tmpExpandedValue, expandErr := expand(activeContext, activeProperty,
					value)
				if !isNil(expandErr) {
					return nil, expandErr
				}
				expandedValue = tmpExpandedValue
				//TODO the step between 7.4.9.2 and 7.4.9.3 is not in the spec
				//but it should definitely be.
				if _, isArray := expandedValue.([]interface{}); !isArray {
					tmpArray := make([]interface{}, 0)
					tmpArray = append(tmpArray, expandedValue)
					expandedValue = tmpArray
				}
				// 7.4.9.3)
				if isListObject(expandedValue) {
					return nil, LIST_OF_LISTS
				}
			}
			// 7.4.10)
			if *expandedProperty == "@set" {
				tmpExpandedValue, expandErr := expand(activeContext, activeProperty,
					value)
				if !isNil(expandErr) {
					return nil, expandErr
				}
				expandedValue = tmpExpandedValue
			}
			// 7.4.11)
			if *expandedProperty == "@reverse" {
				if !isMap {
					return nil, INVALID_REVERSE_VALUE
				}
				// 7.4.11.1)
				reverseArg := "@reverse"
				tmpExpandedValue, expandErr := expand(activeContext, &reverseArg,
					value)
				if !isNil(expandErr) {
					return nil, expandErr
				}
				expandedValue = tmpExpandedValue
				// 7.4.11.2)
				expandedValueMap := expandedValue.(map[string]interface{})
				reverse, hasReverse := expandedValueMap["@reverse"]
				reverseMap, isReverseMap := reverse.(map[string]interface{})
				if hasReverse && isReverseMap {
					for property, item := range reverseMap {
						// 7.4.11.2.1)
						if _, hasProperty := result[property]; !hasProperty {
							result[property] = make([]interface{}, 0)
						}
						// 7.4.11.2.1)
						resultArray := result[property].([]interface{})
						if itemArray, isArray := item.([]interface{}); isArray {
							for _, subItem := range itemArray {
								resultArray = append(resultArray, subItem)
							}
						} else {
							resultArray = append(resultArray, item)
						}
						result[property] = resultArray
					}
				}
				// 7.4.11.3)
				if (!hasReverse && len(expandedValueMap) > 0) ||
					(hasReverse && len(expandedValueMap) > 1) {
					// 7.4.11.3.1)
					if _, hasReverse := result["@reverse"]; !hasReverse {
						result["@reverse"] = make(map[string]interface{})
					}
					// 7.4.11.3.2)
					// Naming the mapping of reverse in result to reverse result instead
					// of reverse map as in the spec because I am already using
					// reverseMap to hold the casting to a map of the variable reverse
					reverseResult := result["@reverse"]
					reverseResultMap := reverseResult.(map[string]interface{})
					// 7.4.11.3.3)
					for property, items := range expandedValueMap {
						if property == "@reverse" {
							continue
						}
						// 7.4.11.3.3.1)
						itemsArray := items.([]interface{})
						for _, item := range itemsArray {
							// 7.4.11.3.3.1.1)
							if isListObject(item) || isValueObject(item) {
								return nil, INVALID_REVERSE_PROPERTY_VALUE
							}
							// 7.4.11.3.3.1.2)
							_, hasProperty := reverseResultMap[property]
							if !hasProperty {
								reverseResultMap[property] = make([]interface{}, 0)
							}
							// 7.4.11.3.3.1.3)
							reverseArray := reverseResultMap[property].([]interface{})
							reverseResultMap[property] = append(reverseArray, item)
						}
					}
				}
				// 7.4.11.4)
				continue
			}
			//TODO java code differs from spec here
			// 7.4.12)
			if !isNil(expandedValue) {
				result[*expandedProperty] = expandedValue
			}
			// 7.4.13)
			continue
			// 7.5)
		} else if _, isValueMap := value.(map[string]interface{}); isValueMap &&
			activeContext.getContainer(key) == "@language" {
			// 7.5.1)
			valueMap := value.(map[string]interface{})
			expandedValue = make([]interface{}, 0)
			// 7.5.2)
			keys := sortedKeys(valueMap)
			for _, language := range keys {
				languageValue := valueMap[language]
				// 7.5.2.1)
				if _, isArray := languageValue.([]interface{}); !isArray {
					tmpArray := make([]interface{}, 0)
					tmpArray = append(tmpArray, languageValue)
					languageValue = tmpArray
				}
				// 7.5.2.2)
				languageArray := languageValue.([]interface{})
				for _, item := range languageArray {
					if _, isString := item.(string); !isString {
						return nil, INVALID_LANGUAGE_MAP_VALUE
					}
					newLanguageMap := make(map[string]interface{})
					newLanguageMap["@language"] = strings.ToLower(language)
					newLanguageMap["@value"] = item
					expandedValue = append(expandedValue.([]interface{}),
						newLanguageMap)
				}
			}
			// 7.6)
		} else if _, isValueMap := value.(map[string]interface{}); isValueMap &&
			activeContext.getContainer(key) == "@index" {
			// 7.1.6)
			valueMap := value.(map[string]interface{})
			expandedValue = make([]interface{}, 0)
			// 7.6.2)
			keys := sortedKeys(valueMap)
			for _, index := range keys {
				indexValue := valueMap[index]
				// 7.6.2.1)
				if _, isArray := indexValue.([]interface{}); !isArray {
					tmpArray := make([]interface{}, 0)
					tmpArray = append(tmpArray, indexValue)
					indexValue = tmpArray
				}
				// 7.6.2.2)
				tmpIndexValue, expandErr := expand(activeContext, &key, indexValue)
				if !isNil(expandErr) {
					return nil, expandErr
				}
				indexValue = tmpIndexValue
				// 7.6.2.3)
				indexArray := indexValue.([]interface{})
				for _, item := range indexArray {
					// 7.6.2.3.1)
					itemMap := item.(map[string]interface{})
					if _, hasIndex := itemMap["@index"]; !hasIndex {
						itemMap["@index"] = index
					}
					// 7.6.2.3.2)
					expandedValue = append(expandedValue.([]interface{}), item)
				}
			}
		} else {
			// 7.7)
			tmpExpandedValue, expandErr := expand(activeContext, &key, value)
			if !isNil(expandErr) {
				return nil, expandErr
			}
			expandedValue = tmpExpandedValue
		}
		// 7.8)
		if isNil(expandedValue) {
			continue
		}
		// 7.9)
		if !isListObject(expandedValue) &&
			"@list" == activeContext.getContainer(key) {
			if _, isValueArray := expandedValue.([]interface{}); !isValueArray {
				tmpArray := make([]interface{}, 0)
				tmpArray = append(tmpArray, expandedValue)
				expandedValue = tmpArray
			}
			tmpMap := make(map[string]interface{}, 0)
			tmpMap["@list"] = expandedValue
			expandedValue = tmpMap
		}
		//7.10)
		if activeContext.isReverseProperty(key) {
			// 7.10.1)
			if _, hasReverse := result["@reverse"]; !hasReverse {
				result["@reverse"] = make(map[string]interface{})
			}
			// 7.10.2)
			reverseMap := result["@reverse"].(map[string]interface{})
			// 7.10.3)
			if _, isExpandedArray := expandedValue.([]interface{}); !isExpandedArray {
				tmpArray := make([]interface{}, 0)
				tmpArray = append(tmpArray, expandedValue)
				expandedValue = tmpArray
			}
			// 7.10.4)
			expandedArray := expandedValue.([]interface{})
			for _, item := range expandedArray {
				// 7.10.4.1)
				if isValueObject(item) || isListObject(item) {
					return nil, INVALID_REVERSE_PROPERTY_VALUE
				}
				// 7.10.4.2)
				_, hasProperty := reverseMap[*expandedProperty]
				if !hasProperty {
					reverseMap[*expandedProperty] = make([]interface{}, 0)
				}
				// 7.10.4.3)
				reverseArray := reverseMap[*expandedProperty].([]interface{})
				if itemArray, isArray := item.([]interface{}); isArray {
					for _, subItem := range itemArray {
						reverseArray = append(reverseArray, subItem)
					}
				} else {
					reverseArray = append(reverseArray, item)
				}
				reverseMap[*expandedProperty] = reverseArray
			}
			// 7.11)
		} else {
			// 7.11.1)
			if _, hasProperty := result[*expandedProperty]; !hasProperty {
				result[*expandedProperty] = make([]interface{}, 0)
			}
			// 7.11.2
			resultArray := result[*expandedProperty].([]interface{})
			if expandedArray, isArray := expandedValue.([]interface{}); isArray {
				for _, item := range expandedArray {
					resultArray = append(resultArray, item)
				}
			} else {
				resultArray = append(resultArray, expandedValue)
			}
			result[*expandedProperty] = resultArray
		}
	}
	// 8)
	if value, hasValue := result["@value"]; hasValue {
		//8.1)
		if !isValidValueObject(result) {
			return nil, INVALID_VALUE_OBJECT
		}
		// 8.2)
		if isNil(value) {
			result = nil
			// 8.3)
		} else if _, isValueString := value.(string); !isValueString {
			if _, hasLanguage := result["@language"]; hasLanguage {
				return nil, INVALID_LANGUAGE_TAGGED_VALUE
			}
		} else if typeVal, hasType := result["@type"]; hasType {
			// 8.4)
			if !isIRI(typeVal) {
				return nil, INVALID_TYPED_VALUE
			}
		}
	} else if typeVal, hasType := result["@type"]; hasType {
		// 9)
		if _, isTypeArray := typeVal.([]interface{}); !isTypeArray {
			tmpArray := make([]interface{}, 0)
			tmpArray = append(tmpArray, typeVal)
			result["@type"] = tmpArray
		}
	} else {
		// 10)
		_, hasSet := result["@set"]
		_, hasList := result["@list"]
		if hasSet || hasList {
			// 10.1)
			maxLen := 0
			if _, hasIndex := result["@index"]; hasIndex {
				maxLen = 2
			} else {
				maxLen = 1
			}
			if len(result) > maxLen {
				return nil, INVALID_SET_OR_LIST_OBJECT
			}
			// 10.2)
			if hasSet {
				// TODO check comment's validity
				// result becomes an array here, thus the remaining checks
				// will never be true from here on
				// so simply return the value rather than have to make
				// result an object and cast it with every
				// other use in the function.
				return result["@set"], nil
			}
		}
	}
	// 11)
	if _, hasLanguage := result["@language"]; hasLanguage &&
		len(result) == 1 {
		result = nil
	}
	// 12)
	if isNil(activeProperty) || *activeProperty == "@graph" {
		// 12.1)
		_, hasValue := result["@value"]
		_, hasList := result["@list"]
		_, hasID := result["@id"]
		if !isNil(result) && (len(result) == 0 || hasList || hasValue) {
			result = nil
		} else if !isNil(result) && len(result) == 1 && hasID {
			// 12.2)
			result = nil
		}
	}
	// 13)
	return result, nil
}

func expandValue(activeContext *Context, activeProperty string,
	value interface{}) (interface{}, error) {
	result := make(map[string]interface{})
	termDefinition, hasDefinition := activeContext.termDefinitions[activeProperty]
	termMap, _ := termDefinition.(map[string]interface{})
	typeValue, hasType := termMap["@type"]
	// 1)
	if hasDefinition && hasType && typeValue == "@id" {
		valueString := value.(string)
		expandedValue, expandErr := expandIri(activeContext, &valueString,
			true, false, nil, nil)
		if isNil(expandErr) {
			result["@id"] = *expandedValue
			return result, nil
		} else {
			return nil, expandErr
		}
	}
	// 2)
	if hasDefinition && hasType && typeValue == "@vocab" {
		valueString := value.(string)
		expandedValue, expandErr := expandIri(activeContext, &valueString,
			true, true, nil, nil)
		if isNil(expandErr) {
			result["@id"] = *expandedValue
			return result, nil
		} else {
			return nil, expandErr
		}
	}
	// 3)
	result["@value"] = value
	// 4)
	if hasDefinition && hasType {
		result["@type"] = typeValue
	} else if _, isString := value.(string); isString {
		// 5.1)
		language, hasLanguage := termMap["@language"]
		defaultLanguage, hasDefaultLanguage := activeContext.table["@language"]
		if hasDefinition && hasLanguage {
			if !isNil(language) {
				result["@language"] = language
			}
			// 5.2)
		} else if hasDefaultLanguage {
			result["@language"] = defaultLanguage
		}
	}
	// 6)
	return result, nil
}
