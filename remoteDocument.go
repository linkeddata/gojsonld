package gojsonld

type RemoteDocument struct {
	documentUrl string
	document    interface{}
	contextUrl  string
}

func NewRemoteDocument(url string, document interface{}) *RemoteDocument {
	return &RemoteDocument{
		documentUrl: url,
		document:    document,
	}
}

func NewRemoteDocumentContext(url string, document interface{}, context string) *RemoteDocument {
	return &RemoteDocument{
		documentUrl: url,
		document:    document,
		contextUrl:  context,
	}
}
