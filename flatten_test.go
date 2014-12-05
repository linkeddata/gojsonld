package gojsonld

import (
	"encoding/json"
	"reflect"
	"testing"
)

func testFlatten(inputFile, outputFile, contextFile, base string,
	compactArrays bool, t *testing.T) {
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
	contextJson, jsonErr := ReadJSONFromFile(test_dir + contextFile)
	if !isNil(jsonErr) {
		contextJson = nil
	}
	options := &Options{
		Base:           base,
		CompactArrays:  compactArrays,
		ExpandContext:  nil,
		DocumentLoader: NewDocumentLoader(),
	}
	flattenedJson, flattenErr := Flatten(inputJson, contextJson, options)
	if !isNil(flattenErr) {
		t.Error("Compaction failed with error ", flattenErr.Error())
		return
	}
	flattenedString, _ := json.MarshalIndent(flattenedJson, "", "    ")
	outputString, _ := json.MarshalIndent(outputJson, "", "    ")
	if !reflect.DeepEqual(flattenedJson, outputJson) {
		t.Error("Expected:\n", string(outputString), "\nGot:\n",
			string(flattenedString))
	}
}

func testFlattenSimple(inputFile, outputFile, base string,
	t *testing.T) {
	testFlatten(inputFile, outputFile, "", base, true, t)
}

func TestFlatten0001(t *testing.T) {
	testFlattenSimple("flatten-0001-in.jsonld", "flatten-0001-out.jsonld",
		"", t)
}

func TestFlatten0002(t *testing.T) {
	testFlattenSimple("flatten-0002-in.jsonld", "flatten-0002-out.jsonld",
		"", t)
}

func TestFlatten0003(t *testing.T) {
	testFlattenSimple("flatten-0003-in.jsonld", "flatten-0003-out.jsonld",
		"", t)
}

func TestFlatten0004(t *testing.T) {
	testFlattenSimple("flatten-0004-in.jsonld", "flatten-0004-out.jsonld",
		"", t)
}

func TestFlatten0005(t *testing.T) {
	testFlattenSimple("flatten-0005-in.jsonld", "flatten-0005-out.jsonld",
		"http://json-ld.org/test-suite/tests/flatten-0005-in.jsonld", t)
}

func TestFlatten0006(t *testing.T) {
	testFlattenSimple("flatten-0006-in.jsonld", "flatten-0006-out.jsonld",
		"", t)
}

func TestFlatten0007(t *testing.T) {
	testFlattenSimple("flatten-0007-in.jsonld", "flatten-0007-out.jsonld",
		"", t)
}

func TestFlatten0008(t *testing.T) {
	testFlattenSimple("flatten-0008-in.jsonld", "flatten-0008-out.jsonld",
		"", t)
}

func TestFlatten0009(t *testing.T) {
	testFlattenSimple("flatten-0009-in.jsonld", "flatten-0009-out.jsonld",
		"", t)
}

func TestFlatten0010(t *testing.T) {
	testFlattenSimple("flatten-0010-in.jsonld", "flatten-0010-out.jsonld",
		"", t)
}

func TestFlatten0011(t *testing.T) {
	testFlattenSimple("flatten-0011-in.jsonld", "flatten-0011-out.jsonld",
		"", t)
}

func TestFlatten0012(t *testing.T) {
	testFlattenSimple("flatten-0012-in.jsonld", "flatten-0012-out.jsonld",
		"", t)
}

func TestFlatten0013(t *testing.T) {
	testFlattenSimple("flatten-0013-in.jsonld", "flatten-0013-out.jsonld",
		"", t)
}

func TestFlatten0014(t *testing.T) {
	testFlattenSimple("flatten-0014-in.jsonld", "flatten-0014-out.jsonld",
		"", t)
}

func TestFlatten0015(t *testing.T) {
	testFlattenSimple("flatten-0015-in.jsonld", "flatten-0015-out.jsonld",
		"", t)
}

func TestFlatten0016(t *testing.T) {
	testFlattenSimple("flatten-0016-in.jsonld", "flatten-0016-out.jsonld",
		"", t)
}

func TestFlatten0017(t *testing.T) {
	testFlattenSimple("flatten-0017-in.jsonld", "flatten-0017-out.jsonld",
		"", t)
}

func TestFlatten0018(t *testing.T) {
	testFlattenSimple("flatten-0018-in.jsonld", "flatten-0018-out.jsonld",
		"", t)
}

func TestFlatten0019(t *testing.T) {
	testFlattenSimple("flatten-0019-in.jsonld", "flatten-0019-out.jsonld",
		"", t)
}

func TestFlatten0020(t *testing.T) {
	testFlattenSimple("flatten-0020-in.jsonld", "flatten-0020-out.jsonld",
		"", t)
}

func TestFlatten0021(t *testing.T) {
	testFlattenSimple("flatten-0021-in.jsonld", "flatten-0021-out.jsonld",
		"", t)
}

func TestFlatten0022(t *testing.T) {
	testFlattenSimple("flatten-0022-in.jsonld", "flatten-0022-out.jsonld",
		"", t)
}

func TestFlatten0023(t *testing.T) {
	testFlattenSimple("flatten-0023-in.jsonld", "flatten-0023-out.jsonld",
		"", t)
}

func TestFlatten0024(t *testing.T) {
	testFlattenSimple("flatten-0024-in.jsonld", "flatten-0024-out.jsonld",
		"", t)
}

func TestFlatten0025(t *testing.T) {
	testFlattenSimple("flatten-0025-in.jsonld", "flatten-0025-out.jsonld",
		"", t)
}

func TestFlatten0026(t *testing.T) {
	testFlattenSimple("flatten-0026-in.jsonld", "flatten-0026-out.jsonld",
		"", t)
}

func TestFlatten0027(t *testing.T) {
	testFlattenSimple("flatten-0027-in.jsonld", "flatten-0027-out.jsonld",
		"", t)
}

func TestFlatten0028(t *testing.T) {
	testFlattenSimple("flatten-0028-in.jsonld", "flatten-0028-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestFlatten0029(t *testing.T) {
	testFlattenSimple("flatten-0029-in.jsonld", "flatten-0029-out.jsonld",
		"http://json-ld.org/test-suite/tests/flatten-0029-in.jsonld", t)
}

func TestFlatten0030(t *testing.T) {
	testFlattenSimple("flatten-0030-in.jsonld", "flatten-0030-out.jsonld",
		"", t)
}

func TestFlatten0031(t *testing.T) {
	testFlattenSimple("flatten-0031-in.jsonld", "flatten-0031-out.jsonld",
		"", t)
}

func TestFlatten0032(t *testing.T) {
	testFlattenSimple("flatten-0032-in.jsonld", "flatten-0032-out.jsonld",
		"", t)
}

func TestFlatten0033(t *testing.T) {
	testFlattenSimple("flatten-0033-in.jsonld", "flatten-0033-out.jsonld",
		"", t)
}

func TestFlatten0034(t *testing.T) {
	testFlattenSimple("flatten-0034-in.jsonld", "flatten-0034-out.jsonld",
		"", t)
}

func TestFlatten0035(t *testing.T) {
	testFlattenSimple("flatten-0035-in.jsonld", "flatten-0035-out.jsonld",
		"", t)
}

func TestFlatten0036(t *testing.T) {
	testFlattenSimple("flatten-0036-in.jsonld", "flatten-0036-out.jsonld",
		"", t)
}

func TestFlatten0037(t *testing.T) {
	testFlattenSimple("flatten-0037-in.jsonld", "flatten-0037-out.jsonld",
		"", t)
}

func TestFlatten0038(t *testing.T) {
	testFlattenSimple("flatten-0038-in.jsonld", "flatten-0038-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestFlatten0039(t *testing.T) {
	testFlattenSimple("flatten-0039-in.jsonld", "flatten-0039-out.jsonld",
		"", t)
}

func TestFlatten0040(t *testing.T) {
	testFlattenSimple("flatten-0040-in.jsonld", "flatten-0040-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestFlatten0041(t *testing.T) {
	testFlattenSimple("flatten-0041-in.jsonld", "flatten-0041-out.jsonld",
		"", t)
}

func TestFlatten0042(t *testing.T) {
	testFlattenSimple("flatten-0042-in.jsonld", "flatten-0042-out.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestFlatten0043(t *testing.T) {
	testFlattenSimple("flatten-0043-in.jsonld", "flatten-0043-out.jsonld",
		"http://json-ld.org/test-suite/tests/flatten-0043-in.jsonld", t)
}

func TestFlatten0044(t *testing.T) {
	testFlatten("flatten-0044-in.jsonld", "flatten-0044-out.jsonld",
		"flatten-0044-context.jsonld", "", false, t)
}

func TestFlatten0045(t *testing.T) {
	testFlattenSimple("flatten-0045-in.jsonld", "flatten-0045-out.jsonld",
		"", t)
}

func TestFlatten0046(t *testing.T) {
	testFlattenSimple("flatten-0046-in.jsonld", "flatten-0046-out.jsonld",
		"", t)
}
