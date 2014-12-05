package gojsonld

import (
	"strings"
)

func parse(activeContext *Context, localContext interface{},
	remoteContexts []string) (*Context, error) {
	if isNil(remoteContexts) {
		remoteContexts = make([]string, 0)
	}
	// 1)
	result := activeContext.clone()
	// 2)
	if _, isArray := localContext.([]interface{}); !isArray {
		tmpArray := make([]interface{}, 0)
		tmpArray = append(tmpArray, localContext)
		localContext = tmpArray
	}
	// 3)
	for _, context := range localContext.([]interface{}) {
		// 3.1)
		//TODO The base IRI of the active context is set
		//to the IRI of the currently being processed
		//document (which might be different from the currently
		//being processed context), if available; otherwise
		//to null. If set, the base option of a JSON-LD API
		//Implementation overrides the base IRI.
		if isNil(context) {
			newContext := new(Context)
			newContext.init(activeContext.options)
			result = newContext
			continue
		}
		// 3.2)
		if contextString, isString := context.(string); isString {
			// 3.2.1)
			uri := result.table["@base"].(string)
			absoluteUri, resolveErr := resolve(&uri, &contextString)
			if !isNil(resolveErr) {
				return nil, resolveErr
			}
			uri = *absoluteUri
			// 3.2.2
			isRecursive := false
			for _, remoteContext := range remoteContexts {
				if remoteContext == contextString {
					isRecursive = true
					break
				}
			}
			if isRecursive {
				return nil, RECURSIVE_CONTEXT_INCLUSION
			}
			remoteContexts = append(remoteContexts, contextString)
			// 3.2.3
			rd, loadErr := activeContext.options.DocumentLoader.
				loadDocument(contextString)
			if !isNil(loadErr) {
				return nil, loadErr
			}
			var remoteContext interface{} = rd.document
			remoteContextMap, isMap := remoteContext.(map[string]interface{})
			_, containsContext := remoteContextMap["@context"]
			if !isMap {
				return nil, LOADING_REMOTE_CONTEXT_FAILED
			} else if !containsContext {
				return nil, INVALID_REMOTE_CONTEXT
			}
			context = remoteContextMap["@context"]
			// 3.2.4)
			recursiveResult, parseErr := parse(result, context, remoteContexts)
			if isNil(parseErr) {
				result = recursiveResult
			} else {
				return nil, parseErr
			}
			// 3.2.5)
			continue
		}
		// 3.3)
		contextMap, isMap := context.(map[string]interface{})
		if !isMap {
			return nil, INVALID_LOCAL_CONTEXT
		}
		// 3.4)
		// 3.4.1)
		value, hasBase := contextMap["@base"]
		if len(remoteContexts) == 0 && hasBase {
			// 3.4.2)
			if isNil(value) {
				delete(result.table, "@base")
			} else if valueString, isString := value.(string); isString {
				//TODO check isAbsoluteIri
				if isAbsoluteIri(valueString) {
					// 3.4.3)
					result.table["@base"] = valueString
				} else {
					// 3.4.4)
					baseURI := result.table["@base"]
					if !isAbsoluteIri(baseURI.(string)) {
						//3.4.5)
						return nil, INVALID_BASE_IRI
					}
					//TODO
					//result.put("@base", JsonLdUrl.resolve(baseUri, (String) baseValue));
				}
			} else {
				return nil, INVALID_BASE_IRI
			}
		}
		// 3.5)
		// 3.5.1)
		value, hasVocab := contextMap["@vocab"]
		if hasVocab {
			// 3.5.2)
			if isNil(value) {
				delete(result.table, "@vocab")
			} else if valueString, isString := value.(string); isString {
				// 3.5.3)
				if isAbsoluteIri(valueString) || isBlankNodeIdentifier(valueString) {
					result.table["@vocab"] = valueString
				} else {
					return nil, INVALID_VOCAB_MAPPING
				}
			} else {
				return nil, INVALID_VOCAB_MAPPING
			}
		}
		// 3.6)
		// 3.6.1)
		value, hasLanguage := contextMap["@language"]
		if hasLanguage {
			if isNil(value) {
				delete(result.table, "@language")
			} else if valueString, isString := value.(string); isString {
				result.table["@language"] = strings.ToLower(valueString)
			} else {
				return nil, INVALID_DEFAULT_LANGUAGE
			}
		}
		// 3.7
		defined := make(map[string]bool, 0)
		for key := range contextMap {
			if key == "@base" || key == "@vocab" || key == "@language" {
				continue
			}
			termErr := createTermDefinition(result, contextMap, key, defined)
			if !isNil(termErr) {
				return nil, termErr
			}
		}
	}
	return result, nil
}

func createTermDefinition(activeContext *Context, localContext map[string]interface{},
	term string, defined map[string]bool) error {
	// 1)
	if definedValue, isDefined := defined[term]; isDefined {
		if definedValue {
			return nil
		}
		return CYCLIC_IRI_MAPPING
	}
	// 2)
	defined[term] = false
	// 3)
	if isKeyword(term) {
		return KEYWORD_REDEFINITION
	}
	// 4)
	delete(activeContext.termDefinitions, term)
	// 5)
	var value interface{} = deepCopy(localContext[term])
	// 6)
	valueMap, isMap := value.(map[string]interface{})
	idValue, hasId := valueMap["@id"]
	if isNil(value) || (isMap && hasId && isNil(idValue)) {
		activeContext.termDefinitions[term] = nil
		defined[term] = true
		return nil
	}
	// 7)
	if _, isString := value.(string); isString {
		tmpMap := make(map[string]interface{}, 0)
		tmpMap["@id"] = value
		value = tmpMap
		// 8
	}
	if _, isMap := value.(map[string]interface{}); !isMap {
		return INVALID_TERM_DEFINITION
	}
	//Redifine valueMap
	valueMap = value.(map[string]interface{})
	//9)
	definition := make(map[string]interface{}, 0)
	// 10)
	if typeVal, hasType := valueMap["@type"]; hasType {
		// 10.1)
		typeString, isString := typeVal.(string)
		if !isString {
			return INVALID_TYPE_MAPPING
		}
		// 10.2)
		expandedType, expandErr := expandIri(activeContext, &typeString, false,
			true, localContext, defined)
		if isNil(expandErr) {
			typeString = *expandedType
		} else {
			return expandErr
		}
		if typeString != "@id" && typeString != "@vocab" &&
			!isAbsoluteIri(typeString) {
			return INVALID_TYPE_MAPPING
		}
		// 10.3)
		definition["@type"] = typeString
	}
	// 11)
	if reverse, hasReverse := valueMap["@reverse"]; hasReverse {
		// 11.1)
		if _, hasID := valueMap["@id"]; hasID {
			return INVALID_REVERSE_PROPERTY
		}
		reverseString, isString := reverse.(string)
		// 11.2)
		if !isString {
			return INVALID_IRI_MAPPING
		}
		// 11.3)
		expandedReverse, expandErr := expandIri(activeContext, &reverseString,
			false, true, localContext, defined)
		if isNil(expandErr) {
			reverseString = *expandedReverse
		} else {
			return expandErr
		}
		if !isAbsoluteIri(reverseString) && !isBlankNodeIdentifier(reverseString) {
			return INVALID_IRI_MAPPING
		}
		definition["@id"] = reverseString
		// 11.4)
		if container, hasContainer := valueMap["@container"]; hasContainer {
			if container != "@set" && container != "@index" && !isNil(container) {
				return INVALID_REVERSE_PROPERTY
			}
			definition["@container"] = container
		}
		definition["@reverse"] = true
		activeContext.termDefinitions[term] = definition
		defined[term] = true
		return nil
	}
	//12)
	definition["@reverse"] = false
	// 13)
	id, hasID := valueMap["@id"]
	if hasID && id != term {
		idString, isString := id.(string)
		// 13.1)
		if !isString {
			return INVALID_IRI_MAPPING
		}
		// 13.2)
		expandedID, expandErr := expandIri(activeContext, &idString, false,
			true, localContext, defined)
		if isNil(expandErr) {
			idString = *expandedID
		} else {
			return expandErr
		}
		if isKeyword(idString) || isBlankNodeIdentifier(idString) ||
			isAbsoluteIri(idString) {
			if idString == "@context" {
				return INVALID_KEYWORD_ALIAS
			}
			definition["@id"] = idString
		} else {
			return INVALID_IRI_MAPPING
		}
	} else if strings.Contains(term, ":") {
		// 14)
		colIndex := strings.Index(term, ":")
		prefix := term[:colIndex]
		suffix := term[colIndex+1:]
		// 14.1)
		if _, hasPrefix := localContext[prefix]; hasPrefix {
			createTermDefinition(activeContext, localContext, prefix, defined)
		}
		// 14.2)
		prefixVal, hasPrefix := activeContext.termDefinitions[prefix]
		prefixMap, _ := prefixVal.(map[string]interface{})
		if hasPrefix {
			definition["@id"] = prefixMap["@id"].(string) + suffix
		} else {
			// 14.3)
			definition["@id"] = term
		}
		// 15)
	} else if vocab, hasVocab := activeContext.table["@vocab"]; hasVocab {
		definition["@id"] = vocab.(string) + term
	} else {
		return INVALID_IRI_MAPPING
	}
	// 16)
	// 16.1)
	if container, hasContainer := valueMap["@container"]; hasContainer {
		// 16.2)
		if container != "@list" && container != "@set" && container != "@index" &&
			container != "@language" {
			return INVALID_CONTAINER_MAPPING
		}
		// 16.3)
		definition["@container"] = container
	}
	// 17)
	language, containsLanguage := valueMap["@language"]
	_, containsType := valueMap["@type"]
	if containsLanguage && !containsType {
		// 17.1)
		languageString, isString := language.(string)
		if !isString && !isNil(language) {
			return INVALID_LANGUAGE_MAPPING
		}
		// 17.2)
		if isString {
			definition["@language"] = strings.ToLower(languageString)
		} else {
			definition["@language"] = nil
		}
	}
	//18
	activeContext.termDefinitions[term] = definition
	defined[term] = true
	return nil
}

func expandIri(activeContext *Context, value *string, relative bool, vocab bool,
	localContext map[string]interface{}, defined map[string]bool) (*string, error) {
	//1)
	if isKeyword(*value) || isNil(value) {
		if isNil(value) {
			return nil, nil
		}
		returnValue := *value
		return &returnValue, nil
	}
	//2)
	//TODO figure out what to do when value not in defined
	//for now we take the if branch if value not in defined
	//same thing as in step 4.3
	_, hasValue := localContext[*value]
	definedValue := defined[*value]
	if !isNil(localContext) && hasValue && definedValue == false {
		createErr := createTermDefinition(activeContext, localContext,
			*value, defined)
		if !isNil(createErr) {
			return nil, createErr
		}
	}
	// 3)
	td, hasTermDefinition := activeContext.termDefinitions[*value]
	if vocab && hasTermDefinition {
		if !isNil(td) {
			tdMap := td.(map[string]interface{})
			returnValue := tdMap["@id"].(string)
			return &returnValue, nil
		} else {
			return nil, nil
		}
	}
	// 4)
	if colIndex := strings.Index(*value, ":"); colIndex >= 0 {
		// 4.1)
		prefix := (*value)[:colIndex]
		suffix := (*value)[colIndex+1:]
		// 4.2)
		if prefix == "_" || strings.HasPrefix(suffix, "//") {
			returnValue := *value
			return &returnValue, nil
		}
		// 4.3)
		_, containsPrefix := localContext[prefix]
		definedPrefix, inDefined := defined[prefix]
		if containsPrefix && (!inDefined || definedPrefix == false) {
			createErr := createTermDefinition(activeContext, localContext,
				prefix, defined)
			if !isNil(createErr) {
				return nil, createErr
			}
		}
		// 4.4)
		td, hasTermDefinition := activeContext.termDefinitions[prefix]
		if hasTermDefinition {
			tdMap := td.(map[string]interface{})
			id := tdMap["@id"].(string)
			returnValue := id + suffix
			return &returnValue, nil
		}
		// 4.5
		return value, nil
	}
	// 5)
	vocabMapping, hasVocab := activeContext.table["@vocab"]
	if vocab && hasVocab {
		returnValue := vocabMapping.(string) + *value
		return &returnValue, nil
	} else if relative {
		// 6)
		base, isString := activeContext.table["@base"].(string)
		var baseArg *string
		if isString {
			baseArg = &base
		} else {
			baseArg = nil
		}
		absoluteUri, resolveErr := resolve(baseArg,
			value)
		if !isNil(resolveErr) {
			return nil, resolveErr
		}
		return absoluteUri, nil
	}
	// 7)
	returnValue := *value
	return &returnValue, nil
}
