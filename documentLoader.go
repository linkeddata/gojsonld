package gojsonld

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const ACCEPT_HEADER = "application/ld+json, application/json;q=0.9, application/javascript;q=0.5, text/javascript;q=0.5, text/plain;q=0.2, */*;q=0.1"

type DocumentLoader struct {
	httpClient *http.Client
}

func NewDocumentLoader() *DocumentLoader {
	return &DocumentLoader{
		httpClient: &http.Client{},
	}
}

func (dl *DocumentLoader) loadDocument(uri string) (*RemoteDocument, error) {
	req, reqErr := http.NewRequest("GET", uri, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Set("Accept", ACCEPT_HEADER)
	res, resErr := dl.httpClient.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	var jsonBody interface{}
	jsonErr := json.Unmarshal(body, &jsonBody)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return NewRemoteDocument(uri, jsonBody), nil
}
