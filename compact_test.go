package gojsonld

import (
	"encoding/json"
	"reflect"
	"testing"
)

func testCompact(inputFile, outputFile, contextFile, base string,
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
		t.Error("Could not open context file")
		return
	}
	options := &Options{
		Base:           base,
		CompactArrays:  compactArrays,
		ExpandContext:  nil,
		DocumentLoader: NewDocumentLoader(),
	}
	compactedJson, compactErr := Compact(inputJson, contextJson, options)
	if !isNil(compactErr) {
		t.Error("Compaction failed with error ", compactErr.Error())
		return
	}
	compactedString, _ := json.MarshalIndent(compactedJson, "", "    ")
	outputString, _ := json.MarshalIndent(outputJson, "", "    ")
	if !reflect.DeepEqual(compactedJson, outputJson) {
		t.Error("Expected:\n", string(outputString), "\nGot:\n",
			string(compactedString))
	}
}

func testCompactSimple(inputFile, outputFile, contextFile, base string,
	t *testing.T) {
	testCompact(inputFile, outputFile, contextFile, base, true, t)
}

func TestCompact0001(t *testing.T) {
	testCompactSimple("compact-0001-in.jsonld", "compact-0001-out.jsonld",
		"compact-0001-context.jsonld", "", t)
}

func TestCompact0002(t *testing.T) {
	testCompactSimple("compact-0002-in.jsonld", "compact-0002-out.jsonld",
		"compact-0002-context.jsonld", "", t)
}

func TestCompact0003(t *testing.T) {
	testCompactSimple("compact-0003-in.jsonld", "compact-0003-out.jsonld",
		"compact-0003-context.jsonld", "", t)
}

func TestCompact0004(t *testing.T) {
	testCompactSimple("compact-0004-in.jsonld", "compact-0004-out.jsonld",
		"compact-0004-context.jsonld", "", t)
}

func TestCompact0005(t *testing.T) {
	testCompactSimple("compact-0005-in.jsonld", "compact-0005-out.jsonld",
		"compact-0005-context.jsonld", "", t)
}

func TestCompact0006(t *testing.T) {
	testCompactSimple("compact-0006-in.jsonld", "compact-0006-out.jsonld",
		"compact-0006-context.jsonld", "", t)
}

func TestCompact0007(t *testing.T) {
	testCompactSimple("compact-0007-in.jsonld", "compact-0007-out.jsonld",
		"compact-0007-context.jsonld", "", t)
}

func TestCompact0008(t *testing.T) {
	testCompactSimple("compact-0008-in.jsonld", "compact-0008-out.jsonld",
		"compact-0008-context.jsonld", "", t)
}

func TestCompact0009(t *testing.T) {
	testCompactSimple("compact-0009-in.jsonld", "compact-0009-out.jsonld",
		"compact-0009-context.jsonld", "", t)
}

func TestCompact0010(t *testing.T) {
	testCompactSimple("compact-0010-in.jsonld", "compact-0010-out.jsonld",
		"compact-0010-context.jsonld", "", t)
}

func TestCompact0011(t *testing.T) {
	testCompactSimple("compact-0011-in.jsonld", "compact-0011-out.jsonld",
		"compact-0011-context.jsonld", "", t)
}

func TestCompact0012(t *testing.T) {
	testCompactSimple("compact-0012-in.jsonld", "compact-0012-out.jsonld",
		"compact-0012-context.jsonld", "", t)
}

func TestCompact0013(t *testing.T) {
	testCompactSimple("compact-0013-in.jsonld", "compact-0013-out.jsonld",
		"compact-0013-context.jsonld", "", t)
}

func TestCompact0014(t *testing.T) {
	testCompactSimple("compact-0014-in.jsonld", "compact-0014-out.jsonld",
		"compact-0014-context.jsonld", "", t)
}

func TestCompact0015(t *testing.T) {
	testCompactSimple("compact-0015-in.jsonld", "compact-0015-out.jsonld",
		"compact-0015-context.jsonld", "", t)
}

func TestCompact0016(t *testing.T) {
	testCompactSimple("compact-0016-in.jsonld", "compact-0016-out.jsonld",
		"compact-0016-context.jsonld", "", t)
}

func TestCompact0017(t *testing.T) {
	testCompactSimple("compact-0017-in.jsonld", "compact-0017-out.jsonld",
		"compact-0017-context.jsonld", "", t)
}

func TestCompact0018(t *testing.T) {
	testCompactSimple("compact-0018-in.jsonld", "compact-0018-out.jsonld",
		"compact-0018-context.jsonld", "", t)
}

func TestCompact0019(t *testing.T) {
	testCompactSimple("compact-0019-in.jsonld", "compact-0019-out.jsonld",
		"compact-0019-context.jsonld", "", t)
}

func TestCompact0020(t *testing.T) {
	testCompactSimple("compact-0020-in.jsonld", "compact-0020-out.jsonld",
		"compact-0020-context.jsonld", "", t)
}

func TestCompact0021(t *testing.T) {
	testCompactSimple("compact-0021-in.jsonld", "compact-0021-out.jsonld",
		"compact-0021-context.jsonld", "", t)
}

func TestCompact0022(t *testing.T) {
	testCompactSimple("compact-0022-in.jsonld", "compact-0022-out.jsonld",
		"compact-0022-context.jsonld", "", t)
}

func TestCompact0023(t *testing.T) {
	testCompactSimple("compact-0023-in.jsonld", "compact-0023-out.jsonld",
		"compact-0023-context.jsonld", "", t)
}

func TestCompact0024(t *testing.T) {
	testCompactSimple("compact-0024-in.jsonld", "compact-0024-out.jsonld",
		"compact-0024-context.jsonld", "", t)
}

func TestCompact0025(t *testing.T) {
	testCompactSimple("compact-0025-in.jsonld", "compact-0025-out.jsonld",
		"compact-0025-context.jsonld", "", t)
}

func TestCompact0026(t *testing.T) {
	testCompactSimple("compact-0026-in.jsonld", "compact-0026-out.jsonld",
		"compact-0026-context.jsonld", "", t)
}

func TestCompact0027(t *testing.T) {
	testCompactSimple("compact-0027-in.jsonld", "compact-0027-out.jsonld",
		"compact-0027-context.jsonld", "", t)
}

func TestCompact0028(t *testing.T) {
	testCompactSimple("compact-0028-in.jsonld", "compact-0028-out.jsonld",
		"compact-0028-context.jsonld", "", t)
}

func TestCompact0029(t *testing.T) {
	testCompactSimple("compact-0029-in.jsonld", "compact-0029-out.jsonld",
		"compact-0029-context.jsonld", "", t)
}

func TestCompact0030(t *testing.T) {
	testCompactSimple("compact-0030-in.jsonld", "compact-0030-out.jsonld",
		"compact-0030-context.jsonld", "", t)
}

func TestCompact0031(t *testing.T) {
	testCompactSimple("compact-0031-in.jsonld", "compact-0031-out.jsonld",
		"compact-0031-context.jsonld", "", t)
}

func TestCompact0032(t *testing.T) {
	testCompactSimple("compact-0032-in.jsonld", "compact-0032-out.jsonld",
		"compact-0032-context.jsonld", "", t)
}

func TestCompact0033(t *testing.T) {
	testCompactSimple("compact-0033-in.jsonld", "compact-0033-out.jsonld",
		"compact-0033-context.jsonld", "", t)
}

func TestCompact0034(t *testing.T) {
	testCompactSimple("compact-0034-in.jsonld", "compact-0034-out.jsonld",
		"compact-0034-context.jsonld", "", t)
}

func TestCompact0035(t *testing.T) {
	testCompactSimple("compact-0035-in.jsonld", "compact-0035-out.jsonld",
		"compact-0035-context.jsonld", "", t)
}

func TestCompact0036(t *testing.T) {
	testCompactSimple("compact-0036-in.jsonld", "compact-0036-out.jsonld",
		"compact-0036-context.jsonld", "", t)
}

func TestCompact0037(t *testing.T) {
	testCompactSimple("compact-0037-in.jsonld", "compact-0037-out.jsonld",
		"compact-0037-context.jsonld",
		"http://json-ld.org/test-suite/tests/", t)
}

func TestCompact0038(t *testing.T) {
	testCompactSimple("compact-0038-in.jsonld", "compact-0038-out.jsonld",
		"compact-0038-context.jsonld", "", t)
}

func TestCompact0039(t *testing.T) {
	testCompactSimple("compact-0039-in.jsonld", "compact-0039-out.jsonld",
		"compact-0039-context.jsonld", "", t)
}

func TestCompact0040(t *testing.T) {
	testCompactSimple("compact-0040-in.jsonld", "compact-0040-out.jsonld",
		"compact-0040-context.jsonld", "", t)
}

func TestCompact0041(t *testing.T) {
	testCompactSimple("compact-0041-in.jsonld", "compact-0041-out.jsonld",
		"compact-0041-context.jsonld", "", t)
}

func TestCompact0042(t *testing.T) {
	testCompactSimple("compact-0042-in.jsonld", "compact-0042-out.jsonld",
		"compact-0042-context.jsonld", "", t)
}

func TestCompact0043(t *testing.T) {
	testCompactSimple("compact-0043-in.jsonld", "compact-0043-out.jsonld",
		"compact-0043-context.jsonld", "", t)
}

func TestCompact0044(t *testing.T) {
	testCompactSimple("compact-0044-in.jsonld", "compact-0044-out.jsonld",
		"compact-0044-context.jsonld", "", t)
}

func TestCompact0045(t *testing.T) {
	testCompactSimple("compact-0045-in.jsonld", "compact-0045-out.jsonld",
		"compact-0045-context.jsonld", "http://json-ld.org/test-suite/tests/", t)
}

func TestCompact0046(t *testing.T) {
	testCompactSimple("compact-0046-in.jsonld", "compact-0046-out.jsonld",
		"compact-0046-context.jsonld", "", t)
}

func TestCompact0047(t *testing.T) {
	testCompactSimple("compact-0047-in.jsonld", "compact-0047-out.jsonld",
		"compact-0047-context.jsonld", "", t)
}

func TestCompact0048(t *testing.T) {
	testCompactSimple("compact-0048-in.jsonld", "compact-0048-out.jsonld",
		"compact-0048-context.jsonld", "", t)
}

func TestCompact0049(t *testing.T) {
	testCompactSimple("compact-0049-in.jsonld", "compact-0049-out.jsonld",
		"compact-0049-context.jsonld", "", t)
}

func TestCompact0050(t *testing.T) {
	testCompactSimple("compact-0050-in.jsonld", "compact-0050-out.jsonld",
		"compact-0050-context.jsonld", "", t)
}

func TestCompact0051(t *testing.T) {
	testCompactSimple("compact-0051-in.jsonld", "compact-0051-out.jsonld",
		"compact-0051-context.jsonld", "", t)
}

func TestCompact0052(t *testing.T) {
	testCompactSimple("compact-0052-in.jsonld", "compact-0052-out.jsonld",
		"compact-0052-context.jsonld", "", t)
}

func TestCompact0053(t *testing.T) {
	testCompactSimple("compact-0053-in.jsonld", "compact-0053-out.jsonld",
		"compact-0053-context.jsonld", "", t)
}

func TestCompact0054(t *testing.T) {
	testCompactSimple("compact-0054-in.jsonld", "compact-0054-out.jsonld",
		"compact-0054-context.jsonld", "", t)
}

func TestCompact0055(t *testing.T) {
	testCompactSimple("compact-0055-in.jsonld", "compact-0055-out.jsonld",
		"compact-0055-context.jsonld", "", t)
}

func TestCompact0056(t *testing.T) {
	testCompactSimple("compact-0056-in.jsonld", "compact-0056-out.jsonld",
		"compact-0056-context.jsonld", "", t)
}

func TestCompact0057(t *testing.T) {
	testCompactSimple("compact-0057-in.jsonld", "compact-0057-out.jsonld",
		"compact-0057-context.jsonld", "", t)
}

func TestCompact0058(t *testing.T) {
	testCompactSimple("compact-0058-in.jsonld", "compact-0058-out.jsonld",
		"compact-0058-context.jsonld", "", t)
}

func TestCompact0059(t *testing.T) {
	testCompactSimple("compact-0059-in.jsonld", "compact-0059-out.jsonld",
		"compact-0059-context.jsonld", "", t)
}

func TestCompact0060(t *testing.T) {
	testCompactSimple("compact-0060-in.jsonld", "compact-0060-out.jsonld",
		"compact-0060-context.jsonld", "", t)
}

func TestCompact0061(t *testing.T) {
	testCompactSimple("compact-0061-in.jsonld", "compact-0061-out.jsonld",
		"compact-0061-context.jsonld", "", t)
}

func TestCompact0062(t *testing.T) {
	testCompactSimple("compact-0062-in.jsonld", "compact-0062-out.jsonld",
		"compact-0062-context.jsonld", "http://json-ld.org/test-suite/tests/", t)
}

func TestCompact0063(t *testing.T) {
	testCompactSimple("compact-0063-in.jsonld", "compact-0063-out.jsonld",
		"compact-0063-context.jsonld", "", t)
}

func TestCompact0064(t *testing.T) {
	testCompactSimple("compact-0064-in.jsonld", "compact-0064-out.jsonld",
		"compact-0064-context.jsonld", "", t)
}

func TestCompact0065(t *testing.T) {
	testCompactSimple("compact-0065-in.jsonld", "compact-0065-out.jsonld",
		"compact-0065-context.jsonld", "", t)
}

func TestCompact0066(t *testing.T) {
	testCompactSimple("compact-0066-in.jsonld", "compact-0066-out.jsonld",
		"compact-0066-context.jsonld",
		"http://json-ld.org/test-suite/tests/compact-0066-in.jsonld", t)
}

func TestCompact0067(t *testing.T) {
	testCompactSimple("compact-0067-in.jsonld", "compact-0067-out.jsonld",
		"compact-0067-context.jsonld", "", t)
}

func TestCompact0068(t *testing.T) {
	testCompactSimple("compact-0068-in.jsonld", "compact-0068-out.jsonld",
		"compact-0068-context.jsonld", "", t)
}

func TestCompact0069(t *testing.T) {
	testCompactSimple("compact-0069-in.jsonld", "compact-0069-out.jsonld",
		"compact-0069-context.jsonld", "", t)
}

func TestCompact0070(t *testing.T) {
	testCompact("compact-0070-in.jsonld", "compact-0070-out.jsonld",
		"compact-0070-context.jsonld", "", false, t)
}

func TestCompact0071(t *testing.T) {
	testCompactSimple("compact-0071-in.jsonld", "compact-0071-out.jsonld",
		"compact-0071-context.jsonld", "", t)
}

func TestCompact0072(t *testing.T) {
	testCompactSimple("compact-0072-in.jsonld", "compact-0072-out.jsonld",
		"compact-0072-context.jsonld", "", t)
}
