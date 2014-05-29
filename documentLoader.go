package gojsonld

import (
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
