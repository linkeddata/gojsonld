package gojsonld

const (
	processingMode = "json-ld-1.0"
)

type Options struct {
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-base
	base string
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-compactArrays
	compactArrays bool
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-expandContext
	expandContext interface{}
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-documentLoader
	documentLoader *DocumentLoader

	// Frame options : http://json-ld.org/spec/latest/json-ld-framing/
	embed       bool
	explicit    bool
	omitDefault bool

	useRdfType            bool
	useNativeTypes        bool
	produceGeneralizedRdf bool

	//TODO
	format        string
	useNamespaces bool
	outputForm    string
}

func NewOptions(base string) *Options {
	return &Options{
		compactArrays: true,
	}
}
