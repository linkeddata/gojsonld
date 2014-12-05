package gojsonld

import (
	"encoding/json"
	"reflect"
	"testing"
)

const test_dir = "./test_files/"

func testExpand(inputFile, outputFile, contextFile, base string, t *testing.T) {
	inputJson, jsonErr := ReadJSONFromFile(test_dir + inputFile)
	if !isNil(jsonErr) {
		t.Error("Could not open input file")
		return
	}
	outputJson, jsonErr := ReadJSONFromFile(test_dir + outputFile)
	if !isNil(jsonErr) {
		t.Error("Could not open output file")
		return
	}
	var expandContext interface{}
	contextJson, jsonErr := ReadJSONFromFile(test_dir + contextFile)
	if !isNil(jsonErr) {
		expandContext = nil
	} else {
		expandContext = contextJson
	}
	options := &Options{
		Base:           base,
		CompactArrays:  true,
		ExpandContext:  expandContext,
		DocumentLoader: NewDocumentLoader(),
	}
	expandedJson, expandErr := Expand(inputJson, options)
	if expandErr != nil {
		t.Error("Expansion failed with error ", expandErr.Error())
		return
	}
	expandedString, _ := json.MarshalIndent(expandedJson, "", "    ")
	outputString, _ := json.MarshalIndent(outputJson, "", "    ")
	if !reflect.DeepEqual(expandedJson, outputJson) {
		t.Error("Expected:\n", string(outputString), "\nGot:\n",
			string(expandedString))
	}
}

func testExpandBase(inputFile, outputFile, base string, t *testing.T) {
	testExpand(inputFile, outputFile, "", base, t)
}

func testExpandContext(inputFile, outputFile, contextFile string, t *testing.T) {
	testExpand(inputFile, outputFile, contextFile, "", t)
}

func testExpandSimple(inputFile, outputFile string, t *testing.T) {
	testExpand(inputFile, outputFile, "", "", t)
}

func TestExpand0001(t *testing.T) {
	testExpandSimple("expand-0001-in.jsonld", "expand-0001-out.jsonld", t)
}

func TestExpand0002(t *testing.T) {
	testExpandSimple("expand-0002-in.jsonld", "expand-0002-out.jsonld", t)
}

func TestExpand0003(t *testing.T) {
	testExpandSimple("expand-0003-in.jsonld", "expand-0003-out.jsonld", t)
}

func TestExpand0004(t *testing.T) {
	testExpandSimple("expand-0004-in.jsonld", "expand-0004-out.jsonld", t)
}

func TestExpand0005(t *testing.T) {
	testExpandBase("expand-0005-in.jsonld", "expand-0005-out.jsonld",
		"http://json-ld.org/test-suite/tests/expand-0005-in.jsonld", t)
}

func TestExpand0006(t *testing.T) {
	testExpandSimple("expand-0006-in.jsonld", "expand-0006-out.jsonld", t)
}

func TestExpand0007(t *testing.T) {
	testExpandSimple("expand-0007-in.jsonld", "expand-0007-out.jsonld", t)
}

func TestExpand0008(t *testing.T) {
	testExpandSimple("expand-0008-in.jsonld", "expand-0008-out.jsonld", t)
}

func TestExpand0009(t *testing.T) {
	testExpandSimple("expand-0009-in.jsonld", "expand-0009-out.jsonld", t)
}

func TestExpand0010(t *testing.T) {
	testExpandSimple("expand-0010-in.jsonld", "expand-0010-out.jsonld", t)
}

func TestExpand0011(t *testing.T) {
	testExpandSimple("expand-0011-in.jsonld", "expand-0011-out.jsonld", t)
}

func TestExpand0012(t *testing.T) {
	testExpandSimple("expand-0012-in.jsonld", "expand-0012-out.jsonld", t)
}

func TestExpand0013(t *testing.T) {
	testExpandSimple("expand-0013-in.jsonld", "expand-0013-out.jsonld", t)
}

func TestExpand0014(t *testing.T) {
	testExpandSimple("expand-0014-in.jsonld", "expand-0014-out.jsonld", t)
}

func TestExpand0015(t *testing.T) {
	testExpandSimple("expand-0015-in.jsonld", "expand-0015-out.jsonld", t)
}

func TestExpand0016(t *testing.T) {
	testExpandSimple("expand-0016-in.jsonld", "expand-0016-out.jsonld", t)
}

func TestExpand0017(t *testing.T) {
	testExpandSimple("expand-0017-in.jsonld", "expand-0017-out.jsonld", t)
}

func TestExpand0018(t *testing.T) {
	testExpandSimple("expand-0018-in.jsonld", "expand-0018-out.jsonld", t)
}

func TestExpand0019(t *testing.T) {
	testExpandSimple("expand-0019-in.jsonld", "expand-0019-out.jsonld", t)
}

func TestExpand0020(t *testing.T) {
	testExpandSimple("expand-0020-in.jsonld", "expand-0020-out.jsonld", t)
}

func TestExpand0021(t *testing.T) {
	testExpandSimple("expand-0021-in.jsonld", "expand-0021-out.jsonld", t)
}

func TestExpand0022(t *testing.T) {
	testExpandSimple("expand-0022-in.jsonld", "expand-0022-out.jsonld", t)
}

func TestExpand0023(t *testing.T) {
	testExpandSimple("expand-0023-in.jsonld", "expand-0023-out.jsonld", t)
}

func TestExpand0024(t *testing.T) {
	testExpandSimple("expand-0024-in.jsonld", "expand-0024-out.jsonld", t)
}

func TestExpand0025(t *testing.T) {
	testExpandSimple("expand-0025-in.jsonld", "expand-0025-out.jsonld", t)
}

func TestExpand0026(t *testing.T) {
	testExpandSimple("expand-0026-in.jsonld", "expand-0026-out.jsonld", t)
}

func TestExpand0027(t *testing.T) {
	testExpandSimple("expand-0027-in.jsonld", "expand-0027-out.jsonld", t)
}

func TestExpand0028(t *testing.T) {
	testExpandBase("expand-0028-in.jsonld", "expand-0028-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0029(t *testing.T) {
	testExpandBase("expand-0029-in.jsonld", "expand-0029-out.jsonld",
		"http://json-ld.org/test-suite/tests/expand-0029-in.jsonld", t)
}

func TestExpand0030(t *testing.T) {
	testExpandSimple("expand-0030-in.jsonld", "expand-0030-out.jsonld", t)
}

func TestExpand0031(t *testing.T) {
	testExpandSimple("expand-0031-in.jsonld", "expand-0031-out.jsonld", t)
}

func TestExpand0032(t *testing.T) {
	testExpandSimple("expand-0032-in.jsonld", "expand-0032-out.jsonld", t)
}

func TestExpand0033(t *testing.T) {
	testExpandSimple("expand-0033-in.jsonld", "expand-0033-out.jsonld", t)
}

func TestExpand0034(t *testing.T) {
	testExpandSimple("expand-0034-in.jsonld", "expand-0034-out.jsonld", t)
}

func TestExpand0035(t *testing.T) {
	testExpandSimple("expand-0035-in.jsonld", "expand-0035-out.jsonld", t)
}

func TestExpand0036(t *testing.T) {
	testExpandSimple("expand-0036-in.jsonld", "expand-0036-out.jsonld", t)
}

func TestExpand0037(t *testing.T) {
	testExpandSimple("expand-0037-in.jsonld", "expand-0037-out.jsonld", t)
}

func TestExpand0038(t *testing.T) {
	testExpandBase("expand-0038-in.jsonld", "expand-0038-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0039(t *testing.T) {
	testExpandSimple("expand-0039-in.jsonld", "expand-0039-out.jsonld", t)
}

func TestExpand0040(t *testing.T) {
	testExpandBase("expand-0040-in.jsonld", "expand-0040-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0041(t *testing.T) {
	testExpandSimple("expand-0041-in.jsonld", "expand-0041-out.jsonld", t)
}

func TestExpand0042(t *testing.T) {
	testExpandSimple("expand-0042-in.jsonld", "expand-0042-out.jsonld", t)
}

func TestExpand0043(t *testing.T) {
	testExpandSimple("expand-0043-in.jsonld", "expand-0043-out.jsonld", t)
}

func TestExpand0045(t *testing.T) {
	testExpandSimple("expand-0045-in.jsonld", "expand-0045-out.jsonld", t)
}

func TestExpand0046(t *testing.T) {
	testExpandSimple("expand-0046-in.jsonld", "expand-0046-out.jsonld", t)
}

func TestExpand0047(t *testing.T) {
	testExpandSimple("expand-0047-in.jsonld", "expand-0047-out.jsonld", t)
}

func TestExpand0048(t *testing.T) {
	testExpandBase("expand-0048-in.jsonld", "expand-0048-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0049(t *testing.T) {
	testExpandSimple("expand-0049-in.jsonld", "expand-0049-out.jsonld", t)
}

func TestExpand0050(t *testing.T) {
	testExpandBase("expand-0050-in.jsonld", "expand-0050-out.jsonld",
		"http://json-ld.org", t)
}

func TestExpand0051(t *testing.T) {
	testExpandBase("expand-0051-in.jsonld", "expand-0051-out.jsonld",
		"http://json-ld.org/", t)
}

func TestExpand0052(t *testing.T) {
	testExpandSimple("expand-0052-in.jsonld", "expand-0052-out.jsonld", t)
}

func TestExpand0053(t *testing.T) {
	testExpandSimple("expand-0053-in.jsonld", "expand-0053-out.jsonld", t)
}

func TestExpand0054(t *testing.T) {
	testExpandSimple("expand-0054-in.jsonld", "expand-0054-out.jsonld", t)
}

func TestExpand0055(t *testing.T) {
	testExpandSimple("expand-0055-in.jsonld", "expand-0055-out.jsonld", t)
}

func TestExpand0056(t *testing.T) {
	testExpandBase("expand-0056-in.jsonld", "expand-0056-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0057(t *testing.T) {
	testExpandBase("expand-0057-in.jsonld", "expand-0057-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0058(t *testing.T) {
	testExpandSimple("expand-0058-in.jsonld", "expand-0058-out.jsonld", t)
}

func TestExpand0059(t *testing.T) {
	testExpandBase("expand-0059-in.jsonld", "expand-0059-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestExpand0060(t *testing.T) {
	testExpandBase("expand-0060-in.jsonld", "expand-0060-out.jsonld",
		"http://json-ld.org/test-suite/tests/expand-0060-in.jsonld", t)
}

func TestExpand0061(t *testing.T) {
	testExpandSimple("expand-0061-in.jsonld", "expand-0061-out.jsonld", t)
}

func TestExpand0062(t *testing.T) {
	testExpandSimple("expand-0062-in.jsonld", "expand-0062-out.jsonld", t)
}

func TestExpand0063(t *testing.T) {
	testExpandSimple("expand-0063-in.jsonld", "expand-0063-out.jsonld", t)
}

func TestExpand0064(t *testing.T) {
	testExpandSimple("expand-0064-in.jsonld", "expand-0064-out.jsonld", t)
}

func TestExpand0065(t *testing.T) {
	testExpandSimple("expand-0065-in.jsonld", "expand-0065-out.jsonld", t)
}

func TestExpand0066(t *testing.T) {
	testExpandBase("expand-0066-in.jsonld", "expand-0066-out.jsonld",
		"http://json-ld.org/test-suite/tests/relative-node", t)
}

func TestExpand0067(t *testing.T) {
	testExpandSimple("expand-0067-in.jsonld", "expand-0067-out.jsonld", t)
}

func TestExpand0068(t *testing.T) {
	testExpandSimple("expand-0068-in.jsonld", "expand-0068-out.jsonld", t)
}

func TestExpand0069(t *testing.T) {
	testExpandSimple("expand-0069-in.jsonld", "expand-0069-out.jsonld", t)
}

func TestExpand0070(t *testing.T) {
	testExpandSimple("expand-0070-in.jsonld", "expand-0070-out.jsonld", t)
}

func TestExpand0071(t *testing.T) {
	testExpandSimple("expand-0071-in.jsonld", "expand-0071-out.jsonld", t)
}

func TestExpand0072(t *testing.T) {
	testExpandSimple("expand-0072-in.jsonld", "expand-0072-out.jsonld", t)
}

func TestExpand0073(t *testing.T) {
	testExpandSimple("expand-0073-in.jsonld", "expand-0073-out.jsonld", t)
}

func TestExpand0074(t *testing.T) {
	testExpandSimple("expand-0074-in.jsonld", "expand-0074-out.jsonld", t)
}

func TestExpand0075(t *testing.T) {
	testExpandSimple("expand-0075-in.jsonld", "expand-0075-out.jsonld", t)
}

func TestExpand0076(t *testing.T) {
	testExpandBase("expand-0076-in.jsonld", "expand-0076-out.jsonld",
		"http://example/base/", t)
}

func TestExpand0077(t *testing.T) {
	testExpandContext("expand-0077-in.jsonld", "expand-0077-out.jsonld",
		"expand-0077-context.jsonld", t)
}

//func TestExpand0078(t *testing.T) {
//testExpandSimple("expand-0078-in.jsonld", "expand-0078-out.jsonld", t)
//}
