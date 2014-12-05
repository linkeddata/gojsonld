package gojsonld

import (
	"net/url"
	"path/filepath"
	"strings"
)

func resolve(base, ref *string) (*string, error) {
	if isNil(base) || *base == "" {
		if isNil(ref) {
			return nil, nil
		}
		returnValue := *ref
		return &returnValue, nil
	}
	if isNil(ref) || strings.Trim(*base, "") == "" {
		returnValue := *base
		return &returnValue, nil
	}
	baseUrl, baseErr := url.Parse(*base)
	refUrl, refErr := url.Parse(*ref)
	if !isNil(baseErr) {
		return nil, baseErr
	}
	if !isNil(refErr) {
		return nil, refErr
	}
	resolvedUrl := baseUrl.ResolveReference(refUrl)
	returnValue := resolvedUrl.String()
	return &returnValue, nil
}

func removeBase(base, ref string) (string, error) {
	if base == "" {
		return ref, nil
	}
	baseUrl, baseErr := url.Parse(base)
	refUrl, refErr := url.Parse(ref)
	if baseErr != nil {
		return "", baseErr
	}
	if refErr != nil {
		return "", refErr
	}
	if baseUrl.Host != refUrl.Host || baseUrl.Scheme != refUrl.Scheme {
		return ref, nil
	}
	rel, err := filepath.Rel(baseUrl.Path, refUrl.Path)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(base, "/") {
		if rel == "." {
			rel = ""
		} else if rel == ".." {
			rel = "./"
		} else if rel == "../.." {
			rel = "../"
		} else if strings.HasPrefix(rel, "../") {
			rel = rel[3:]
		}
	}

	refUrl.Host = ""
	refUrl.Scheme = ""
	refUrl.Path = rel
	return refUrl.String(), nil
}
