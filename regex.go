package gojsonld

import (
	"regexp"
)

var (
	HEX = regexp.MustCompile("[0-9A-Fa-f]")

	PN_CHARS_BASE = regexp.MustCompile("[a-zA-Z]|[\u00C0-\u00D6]|[\u00D8-\u00F6]|[\u00F8-\u02FF]|[\u0370-\u037D]|[\u037F-\u1FFF]|" +
		"[\u200C-\u200D]|[\u2070-\u218F]|[\u2C00-\u2FEF]|[\u3001-\uD7FF]|[\uF900-\uFDCF]|[\uFDF0-\uFFFD]")
	PN_CHARS_U = regexp.MustCompile("(?:" + PN_CHARS_BASE.String() + ")" + "|[_]|[:]")
	PN_CHARS   = regexp.MustCompile("(?:" + PN_CHARS_U.String() + ")" + "|[-]|[0-9]|[\u00B7]|[\u0300-\u036F]|[\u203F-\u2040]")

	ECHAR = regexp.MustCompile("\\\\[tbnrf\\\\\"']")
	UCHAR = regexp.MustCompile("\\\\u" + HEX.String() + HEX.String() + HEX.String() + HEX.String() +
		"|\\\\U" + HEX.String() + HEX.String() + HEX.String() + HEX.String() + HEX.String() + HEX.String() + HEX.String() + HEX.String())

	BLANK_NODE_LABEL = regexp.MustCompile("_:(" + "(?:" + PN_CHARS_U.String() + ")" + "|[0-9])((" + "(?:" + PN_CHARS.String() + ")" +
		"|[\\.])*" + "(?:" + PN_CHARS.String() + ")" + ")?")

	STRING_LITERAL_QUOTE             = regexp.MustCompile("\"([^\u0022\u005C\u000A\u000D]|" + ECHAR.String() + "|" + UCHAR.String() + ")*\"")
	STRING_LITERAL_SINGLE_QUOTE      = regexp.MustCompile("'([^\u0027\u005C\u000A\u000D]|" + ECHAR.String() + "|" + UCHAR.String() + ")*'")
	STRING_LITERAL_LONG_SINGLE_QUOTE = regexp.MustCompile("'''(?:(?:(?:'|'')?[^'\\\\])|" + ECHAR.String() + "|" + UCHAR.String() + ")*'''")
	STRING_LITERAL_LONG_QUOTE        = regexp.MustCompile("\"\"\"(?:(?:(?:\"|\"\")?[^\\\"\\\\])|" + ECHAR.String() + "|" + UCHAR.String() + ")*\"\"\"")

	IRIREF   = regexp.MustCompile("<([^\u0000-\u0020<>\"\\{\\}\\|\\^`\\\\]|" + UCHAR.String() + ")*" + ">")
	EOL      = regexp.MustCompile("[\u000D\u000A]+")
	LANGTAG  = regexp.MustCompile("@[a-zA-Z]+(-[a-zA-Z0-9]+)*")
	DATATYPE = regexp.MustCompile("\\^\\^" + IRIREF.String())
	LITERAL  = regexp.MustCompile("(?:" + STRING_LITERAL_QUOTE.String() + "){1}" + "((?:" + DATATYPE.String() + ")" +
		"|" + "(?:" + LANGTAG.String() + ")" + ")?")
	GRAPH_LABEL = regexp.MustCompile("(?:" + IRIREF.String() + ")" + "|" + "(?:" + BLANK_NODE_LABEL.String() + ")")
	OBJECT      = regexp.MustCompile("(?:" + IRIREF.String() + ")" + "|" + "(?:" + BLANK_NODE_LABEL.String() + ")" +
		"|" + "(?:" + LITERAL.String() + ")")
	//TODO predicate needs to be IRIFEF | BLANK_NODE in order to pass the tests
	PREDICATE = regexp.MustCompile("(?:" + IRIREF.String() + ")" + "|" + "(?:" + BLANK_NODE_LABEL.String() + ")")
	SUBJECT   = regexp.MustCompile("(?:" + IRIREF.String() + ")" + "|" + "(?:" + BLANK_NODE_LABEL.String() + ")")
	STATEMENT = regexp.MustCompile("(?P<subject>" + SUBJECT.String() + ")" + WS_1_N.String() +
		"(?P<predicate>" + PREDICATE.String() + ")" + WS_1_N.String() +
		"(?P<object>" + OBJECT.String() + ")" + WS_1_N.String() +
		"((?P<graph>" + GRAPH_LABEL.String() + ")" + WS_1_N.String() + ")?" + "(?:\\.){1}")

	// TRICKY_UTF_CHARS = regexp.MustCompile("[\uD800\uDC00-\uDB7F\uDFFF]")
	PN_PREFIX = regexp.MustCompile("(?:(?:" + PN_CHARS_BASE.String() + ")(?:(?:" +
		PN_CHARS.String() + "|[\\.])*(?:" + PN_CHARS.String() + "))?)")
	PN_LOCAL_ESC = regexp.MustCompile("[\\\\][_~\\.\\-!$&'\\(\\)*+,;=/?#@%]")
	PERCENT      = regexp.MustCompile("%" + HEX.String() + HEX.String())
	PLX          = regexp.MustCompile(PERCENT.String() + "|" + PN_LOCAL_ESC.String())
	PN_LOCAL     = regexp.MustCompile("((?:" + PN_CHARS_U.String() + "|[:]|[0-9]|" +
		PLX.String() + ")(?:(?:" + PN_CHARS.String() + "|[.]|[:]|" + PLX.String() + ")*(?:" + PN_CHARS.String() + "|[:]|" + PLX.String() +
		"))?)")
	PNAME_NS = regexp.MustCompile("((?:" + PN_PREFIX.String() + ")?):")
	PNAME_LN = regexp.MustCompile("" + PNAME_NS.String() + PN_LOCAL.String())
	WS       = regexp.MustCompile("[ \t]")
	WS_0_N   = regexp.MustCompile(WS.String() + "*")
	WS_0_1   = regexp.MustCompile(WS.String() + "?")
	WS_1_N   = regexp.MustCompile(WS.String() + "+")
	INTEGER  = regexp.MustCompile("[+-]?[0-9]+")
	DECIMAL  = regexp.MustCompile("[+-]?[0-9]*\\.[0-9]+")
	EXPONENT = regexp.MustCompile("[eE][+-]?[0-9]+")
	DOUBLE   = regexp.MustCompile("[+-]?(?:(?:[0-9]+\\.[0-9]*" + EXPONENT.String() +
		")|(?:\\.[0-9]+" + EXPONENT.String() + ")|(?:[0-9]+" + EXPONENT.String() + "))")
)
