package gojsonld

const (
	RDF_SYNTAX_NS = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	RDF_SCHEMA_NS = "http://www.w3.org/2000/01/rdf-schema#"
	XSD_NS        = "http://www.w3.org/2001/XMLSchema#"

	XSD_ANYTYPE = XSD_NS + "anyType"
	XSD_BOOLEAN = XSD_NS + "boolean"
	XSD_DOUBLE  = XSD_NS + "double"
	XSD_INTEGER = XSD_NS + "integer"
	XSD_FLOAT   = XSD_NS + "float"
	XSD_DECIMAL = XSD_NS + "decimal"
	XSD_ANYURI  = XSD_NS + "anyURI"
	XSD_STRING  = XSD_NS + "string"

	RDF_TYPE          = RDF_SYNTAX_NS + "type"
	RDF_FIRST         = RDF_SYNTAX_NS + "first"
	RDF_REST          = RDF_SYNTAX_NS + "rest"
	RDF_NIL           = RDF_SYNTAX_NS + "nil"
	RDF_PLAIN_LITERAL = RDF_SYNTAX_NS + "PlainLiteral"
	RDF_XML_LITERAL   = RDF_SYNTAX_NS + "XMLLiteral"
	RDF_OBJECT        = RDF_SYNTAX_NS + "object"
	RDF_LANGSTRING    = RDF_SYNTAX_NS + "langString"
	RDF_LIST          = RDF_SYNTAX_NS + "List"
)
