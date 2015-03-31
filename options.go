package gojsonld

const (
	processingMode = "json-ld-1.0"
)

type Options struct {
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-base
	Base string
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-compactArrays
	CompactArrays bool
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-expandContext
	ExpandContext interface{}
	// http://www.w3.org/TR/json-ld-api/#widl-JsonLdOptions-documentLoader
	DocumentLoader *DocumentLoader

	// Frame options : http://json-ld.org/spec/latest/json-ld-framing/
	Embed       bool
	Explicit    bool
	OmitDefault bool

	UseRdfType            bool
	UseNativeTypes        bool
	ProduceGeneralizedRdf bool

	//TODO
	Format        string
	UseNamespaces bool
	OutputForm    string
}

func NewOptions(base string) *Options {
	return &Options{
		Base:          base,
		CompactArrays: true,
	}
}
