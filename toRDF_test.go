package gojsonld

import (
	"encoding/json"
	"testing"
)

func testToRDF(input string, output string, produceGeneralizedRDF bool,
	base string, t *testing.T) {
	inputJSON, inputErr := ReadJSONFromFile(test_dir + input)
	if !isNil(inputErr) {
		t.Error("Could not open input file:" + inputErr.Error())
		return
	}
	outputRDF, outputErr := ReadDatasetFromFile(test_dir + output)
	if !isNil(outputErr) {
		t.Error("Could not open output file: " + outputErr.Error())
		return
	}
	options := &Options{}
	options.Base = base
	options.ProduceGeneralizedRdf = produceGeneralizedRDF
	serializedRDF, serializeErr := ToRDF(inputJSON, options)
	if !isNil(serializeErr) {
		t.Error("Could not serialize to RDF: " + serializeErr.Error())
	}
	serializedString, _ := json.MarshalIndent(serializedRDF, "", "    ")
	outputString, _ := json.MarshalIndent(outputRDF, "", "    ")
	if !serializedRDF.Equal(outputRDF) {
		t.Error("Expected:\n", string(outputString), "\nGot:\n",
			string(serializedString))
		return
	}
}

func TestToRDF0001(t *testing.T) {
	testToRDF("toRdf-0001-in.jsonld", "toRdf-0001-out.nq", false, "", t)
}

func TestToRDF0002(t *testing.T) {
	testToRDF("toRdf-0002-in.jsonld", "toRdf-0002-out.nq", false, "", t)
}

func TestToRDF0003(t *testing.T) {
	testToRDF("toRdf-0003-in.jsonld", "toRdf-0003-out.nq", false, "", t)
}

func TestToRDF0004(t *testing.T) {
	testToRDF("toRdf-0004-in.jsonld", "toRdf-0004-out.nq", false, "", t)
}

func TestToRDF0005(t *testing.T) {
	testToRDF("toRdf-0005-in.jsonld", "toRdf-0005-out.nq", false, "", t)
}

func TestToRDF0006(t *testing.T) {
	testToRDF("toRdf-0006-in.jsonld", "toRdf-0006-out.nq", false, "", t)
}

func TestToRDF0007(t *testing.T) {
	testToRDF("toRdf-0007-in.jsonld", "toRdf-0007-out.nq", false, "", t)
}

func TestToRDF0008(t *testing.T) {
	testToRDF("toRdf-0008-in.jsonld", "toRdf-0008-out.nq", false, "", t)
}

func TestToRDF0009(t *testing.T) {
	testToRDF("toRdf-0009-in.jsonld", "toRdf-0009-out.nq", false, "", t)
}

func TestToRDF0010(t *testing.T) {
	testToRDF("toRdf-0010-in.jsonld", "toRdf-0010-out.nq", false, "", t)
}

func TestToRDF0011(t *testing.T) {
	testToRDF("toRdf-0011-in.jsonld", "toRdf-0011-out.nq", false, "", t)
}

func TestToRDF0012(t *testing.T) {
	testToRDF("toRdf-0012-in.jsonld", "toRdf-0012-out.nq", false, "", t)
}

func TestToRDF0013(t *testing.T) {
	testToRDF("toRdf-0013-in.jsonld", "toRdf-0013-out.nq", false, "", t)
}

func TestToRDF0014(t *testing.T) {
	testToRDF("toRdf-0014-in.jsonld", "toRdf-0014-out.nq", false, "", t)
}

func TestToRDF0015(t *testing.T) {
	testToRDF("toRdf-0015-in.jsonld", "toRdf-0015-out.nq", false, "", t)
}

func TestToRDF0016(t *testing.T) {
	testToRDF("toRdf-0016-in.jsonld", "toRdf-0016-out.nq", false,
		"http://json-ld.org/test-suite/tests/toRdf-0016-in.jsonld", t)
}

func TestToRDF0017(t *testing.T) {
	testToRDF("toRdf-0017-in.jsonld", "toRdf-0017-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0018(t *testing.T) {
	testToRDF("toRdf-0018-in.jsonld", "toRdf-0018-out.nq", false,
		"http://json-ld.org/test-suite/tests/toRdf-0018-in.jsonld", t)
}

func TestToRDF0019(t *testing.T) {
	testToRDF("toRdf-0019-in.jsonld", "toRdf-0019-out.nq", false, "", t)
}

func TestToRDF0020(t *testing.T) {
	testToRDF("toRdf-0020-in.jsonld", "toRdf-0020-out.nq", false, "", t)
}

//func TestToRDF0021(t *testing.T) {
//testToRDF("toRdf-0021-in.jsonld", "toRdf-0021-out.nq", false, "", t)
//}

func TestToRDF0022(t *testing.T) {
	testToRDF("toRdf-0022-in.jsonld", "toRdf-0022-out.nq", false, "", t)
}

func TestToRDF0023(t *testing.T) {
	testToRDF("toRdf-0023-in.jsonld", "toRdf-0023-out.nq", false, "", t)
}

func TestToRDF0024(t *testing.T) {
	testToRDF("toRdf-0024-in.jsonld", "toRdf-0024-out.nq", false, "", t)
}

func TestToRDF0025(t *testing.T) {
	testToRDF("toRdf-0025-in.jsonld", "toRdf-0025-out.nq", false, "", t)
}

func TestToRDF0026(t *testing.T) {
	testToRDF("toRdf-0026-in.jsonld", "toRdf-0026-out.nq", false, "", t)
}

func TestToRDF0027(t *testing.T) {
	testToRDF("toRdf-0027-in.jsonld", "toRdf-0027-out.nq", false, "", t)
}

func TestToRDF0028(t *testing.T) {
	testToRDF("toRdf-0028-in.jsonld", "toRdf-0028-out.nq", false, "", t)
}

func TestToRDF0029(t *testing.T) {
	testToRDF("toRdf-0029-in.jsonld", "toRdf-0029-out.nq", false, "", t)
}

func TestToRDF0030(t *testing.T) {
	testToRDF("toRdf-0030-in.jsonld", "toRdf-0030-out.nq", false, "", t)
}

func TestToRDF0031(t *testing.T) {
	testToRDF("toRdf-0031-in.jsonld", "toRdf-0031-out.nq", false, "", t)
}

func TestToRDF0032(t *testing.T) {
	testToRDF("toRdf-0032-in.jsonld", "toRdf-0032-out.nq", false, "", t)
}

func TestToRDF0033(t *testing.T) {
	testToRDF("toRdf-0033-in.jsonld", "toRdf-0033-out.nq", false, "", t)
}

func TestToRDF0034(t *testing.T) {
	testToRDF("toRdf-0034-in.jsonld", "toRdf-0034-out.nq", false, "", t)
}

func TestToRDF0035(t *testing.T) {
	testToRDF("toRdf-0035-in.jsonld", "toRdf-0035-out.nq", false, "", t)
}

func TestToRDF0036(t *testing.T) {
	testToRDF("toRdf-0036-in.jsonld", "toRdf-0036-out.nq", false, "", t)
}

//func TestToRDF0037(t *testing.T) {
//testToRDF("toRdf-0037-in.jsonld", "toRdf-0037-out.nq", false, "", t)
//}

//func TestToRDF0038(t *testing.T) {
//testToRDF("toRdf-0038-in.jsonld", "toRdf-0038-out.nq", false, "", t)
//}

//func TestToRDF0039(t *testing.T) {
//testToRDF("toRdf-0039-in.jsonld", "toRdf-0039-out.nq", false, "", t)
//}

//func TestToRDF0040(t *testing.T) {
//testToRDF("toRdf-0040-in.jsonld", "toRdf-0040-out.nq", false, "", t)
//}

func TestToRDF0041(t *testing.T) {
	testToRDF("toRdf-0041-in.jsonld", "toRdf-0041-out.nq", false, "", t)
}

func TestToRDF0042(t *testing.T) {
	testToRDF("toRdf-0042-in.jsonld", "toRdf-0042-out.nq", false, "", t)
}

func TestToRDF0043(t *testing.T) {
	testToRDF("toRdf-0043-in.jsonld", "toRdf-0043-out.nq", false, "", t)
}

func TestToRDF0044(t *testing.T) {
	testToRDF("toRdf-0044-in.jsonld", "toRdf-0044-out.nq", false, "", t)
}

func TestToRDF0045(t *testing.T) {
	testToRDF("toRdf-0045-in.jsonld", "toRdf-0045-out.nq", false,
		"http://json-ld.org/test-suite/tests/toRdf-0045-in.jsonld", t)
}

func TestToRDF0046(t *testing.T) {
	testToRDF("toRdf-0046-in.jsonld", "toRdf-0046-out.nq", false, "", t)
}

func TestToRDF0047(t *testing.T) {
	testToRDF("toRdf-0047-in.jsonld", "toRdf-0047-out.nq", false, "", t)
}

func TestToRDF0048(t *testing.T) {
	testToRDF("toRdf-0048-in.jsonld", "toRdf-0048-out.nq", false, "", t)
}

func TestToRDF0049(t *testing.T) {
	testToRDF("toRdf-0049-in.jsonld", "toRdf-0049-out.nq", false, "", t)
}

func TestToRDF0050(t *testing.T) {
	testToRDF("toRdf-0050-in.jsonld", "toRdf-0050-out.nq", false, "", t)
}

func TestToRDF0051(t *testing.T) {
	testToRDF("toRdf-0051-in.jsonld", "toRdf-0051-out.nq", false, "", t)
}

func TestToRDF0052(t *testing.T) {
	testToRDF("toRdf-0052-in.jsonld", "toRdf-0052-out.nq", false, "", t)
}

func TestToRDF0053(t *testing.T) {
	testToRDF("toRdf-0053-in.jsonld", "toRdf-0053-out.nq", false, "", t)
}

func TestToRDF0054(t *testing.T) {
	testToRDF("toRdf-0054-in.jsonld", "toRdf-0054-out.nq", false, "", t)
}

func TestToRDF0055(t *testing.T) {
	testToRDF("toRdf-0055-in.jsonld", "toRdf-0055-out.nq", false, "", t)
}

func TestToRDF0056(t *testing.T) {
	testToRDF("toRdf-0056-in.jsonld", "toRdf-0056-out.nq", false, "", t)
}

func TestToRDF0057(t *testing.T) {
	testToRDF("toRdf-0057-in.jsonld", "toRdf-0057-out.nq", false, "", t)
}

func TestToRDF0058(t *testing.T) {
	testToRDF("toRdf-0058-in.jsonld", "toRdf-0058-out.nq", false, "", t)
}

func TestToRDF0059(t *testing.T) {
	testToRDF("toRdf-0059-in.jsonld", "toRdf-0059-out.nq", false, "", t)
}

func TestToRDF0060(t *testing.T) {
	testToRDF("toRdf-0060-in.jsonld", "toRdf-0060-out.nq", false, "", t)
}

func TestToRDF0061(t *testing.T) {
	testToRDF("toRdf-0061-in.jsonld", "toRdf-0061-out.nq", false, "", t)
}

func TestToRDF0062(t *testing.T) {
	testToRDF("toRdf-0062-in.jsonld", "toRdf-0062-out.nq", false, "", t)
}

func TestToRDF0063(t *testing.T) {
	testToRDF("toRdf-0063-in.jsonld", "toRdf-0063-out.nq", false, "", t)
}

func TestToRDF0064(t *testing.T) {
	testToRDF("toRdf-0064-in.jsonld", "toRdf-0064-out.nq", false, "", t)
}

func TestToRDF0065(t *testing.T) {
	testToRDF("toRdf-0065-in.jsonld", "toRdf-0065-out.nq", false, "", t)
}

func TestToRDF0066(t *testing.T) {
	testToRDF("toRdf-0066-in.jsonld", "toRdf-0066-out.nq", false, "", t)
}

func TestToRDF0067(t *testing.T) {
	testToRDF("toRdf-0067-in.jsonld", "toRdf-0067-out.nq", false, "", t)
}

func TestToRDF0068(t *testing.T) {
	testToRDF("toRdf-0068-in.jsonld", "toRdf-0068-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0069(t *testing.T) {
	testToRDF("toRdf-0069-in.jsonld", "toRdf-0069-out.nq", false,
		"http://json-ld.org/test-suite/tests/toRdf-0069-in.jsonld", t)
}

func TestToRDF0070(t *testing.T) {
	testToRDF("toRdf-0070-in.jsonld", "toRdf-0070-out.nq", false, "", t)
}

func TestToRDF0071(t *testing.T) {
	testToRDF("toRdf-0071-in.jsonld", "toRdf-0071-out.nq", false, "", t)
}

func TestToRDF0072(t *testing.T) {
	testToRDF("toRdf-0072-in.jsonld", "toRdf-0072-out.nq", false, "", t)
}

func TestToRDF0073(t *testing.T) {
	testToRDF("toRdf-0073-in.jsonld", "toRdf-0073-out.nq", false, "", t)
}

func TestToRDF0074(t *testing.T) {
	testToRDF("toRdf-0074-in.jsonld", "toRdf-0074-out.nq", false, "", t)
}

func TestToRDF0075(t *testing.T) {
	testToRDF("toRdf-0075-in.jsonld", "toRdf-0075-out.nq", false, "", t)
}

func TestToRDF0076(t *testing.T) {
	testToRDF("toRdf-0076-in.jsonld", "toRdf-0076-out.nq", false, "", t)
}

func TestToRDF0077(t *testing.T) {
	testToRDF("toRdf-0077-in.jsonld", "toRdf-0077-out.nq", false, "", t)
}

func TestToRDF0078(t *testing.T) {
	testToRDF("toRdf-0078-in.jsonld", "toRdf-0078-out.nq", false, "", t)
}

func TestToRDF0079(t *testing.T) {
	testToRDF("toRdf-0079-in.jsonld", "toRdf-0079-out.nq", false, "", t)
}

func TestToRDF0080(t *testing.T) {
	testToRDF("toRdf-0080-in.jsonld", "toRdf-0080-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0081(t *testing.T) {
	testToRDF("toRdf-0081-in.jsonld", "toRdf-0081-out.nq", false, "", t)
}

func TestToRDF0082(t *testing.T) {
	testToRDF("toRdf-0082-in.jsonld", "toRdf-0082-out.nq", false, "", t)
}

func TestToRDF0083(t *testing.T) {
	testToRDF("toRdf-0083-in.jsonld", "toRdf-0083-out.nq", false, "", t)
}

func TestToRDF0084(t *testing.T) {
	testToRDF("toRdf-0084-in.jsonld", "toRdf-0084-out.nq", false, "", t)
}

func TestToRDF0085(t *testing.T) {
	testToRDF("toRdf-0085-in.jsonld", "toRdf-0085-out.nq", false, "", t)
}

func TestToRDF0086(t *testing.T) {
	testToRDF("toRdf-0086-in.jsonld", "toRdf-0086-out.nq", false, "", t)
}

func TestToRDF0087(t *testing.T) {
	testToRDF("toRdf-0087-in.jsonld", "toRdf-0087-out.nq", false, "", t)
}

func TestToRDF0088(t *testing.T) {
	testToRDF("toRdf-0088-in.jsonld", "toRdf-0088-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0089(t *testing.T) {
	testToRDF("toRdf-0089-in.jsonld", "toRdf-0089-out.nq", false, "", t)
}

func TestToRDF0090(t *testing.T) {
	testToRDF("toRdf-0090-in.jsonld", "toRdf-0090-out.nq", false,
		"http://json-ld.org/", t)
}

func TestToRDF0091(t *testing.T) {
	testToRDF("toRdf-0091-in.jsonld", "toRdf-0091-out.nq", false,
		"http://json-ld.org/", t)
}

func TestToRDF0092(t *testing.T) {
	testToRDF("toRdf-0092-in.jsonld", "toRdf-0092-out.nq", false, "", t)
}

func TestToRDF0093(t *testing.T) {
	testToRDF("toRdf-0093-in.jsonld", "toRdf-0093-out.nq", false, "", t)
}

func TestToRDF0094(t *testing.T) {
	testToRDF("toRdf-0094-in.jsonld", "toRdf-0094-out.nq", false, "", t)
}

func TestToRDF0095(t *testing.T) {
	testToRDF("toRdf-0095-in.jsonld", "toRdf-0095-out.nq", false, "", t)
}

func TestToRDF0096(t *testing.T) {
	testToRDF("toRdf-0096-in.jsonld", "toRdf-0096-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0097(t *testing.T) {
	testToRDF("toRdf-0097-in.jsonld", "toRdf-0097-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0098(t *testing.T) {
	testToRDF("toRdf-0098-in.jsonld", "toRdf-0098-out.nq", false, "", t)
}

func TestToRDF0099(t *testing.T) {
	testToRDF("toRdf-0099-in.jsonld", "toRdf-0099-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0100(t *testing.T) {
	testToRDF("toRdf-0100-in.jsonld", "toRdf-0100-out.nq", false,
		"http://json-ld.org/test-suite/tests/toRdf-0100-in.jsonld", t)
}

func TestToRDF0101(t *testing.T) {
	testToRDF("toRdf-0101-in.jsonld", "toRdf-0101-out.nq", false, "", t)
}

func TestToRDF0102(t *testing.T) {
	testToRDF("toRdf-0102-in.jsonld", "toRdf-0102-out.nq", false, "", t)
}

func TestToRDF0103(t *testing.T) {
	testToRDF("toRdf-0103-in.jsonld", "toRdf-0103-out.nq", false, "", t)
}

func TestToRDF0104(t *testing.T) {
	testToRDF("toRdf-0104-in.jsonld", "toRdf-0104-out.nq", false, "", t)
}

func TestToRDF0105(t *testing.T) {
	testToRDF("toRdf-0105-in.jsonld", "toRdf-0105-out.nq", false, "", t)
}

func TestToRDF0106(t *testing.T) {
	testToRDF("toRdf-0106-in.jsonld", "toRdf-0106-out.nq", false,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0107(t *testing.T) {
	testToRDF("toRdf-0107-in.jsonld", "toRdf-0107-out.nq", false, "", t)
}

func TestToRDF0108(t *testing.T) {
	testToRDF("toRdf-0108-in.jsonld", "toRdf-0108-out.nq", false, "", t)
}

func TestToRDF0109(t *testing.T) {
	testToRDF("toRdf-0109-in.jsonld", "toRdf-0109-out.nq", false, "", t)
}

func TestToRDF0110(t *testing.T) {
	testToRDF("toRdf-0110-in.jsonld", "toRdf-0110-out.nq", false, "", t)
}

func TestToRDF0111(t *testing.T) {
	testToRDF("toRdf-0111-in.jsonld", "toRdf-0111-out.nq", false, "", t)
}

func TestToRDF0112(t *testing.T) {
	testToRDF("toRdf-0112-in.jsonld", "toRdf-0112-out.nq", false, "", t)
}

func TestToRDF0113(t *testing.T) {
	testToRDF("toRdf-0113-in.jsonld", "toRdf-0113-out.nq", false, "", t)
}

func TestToRDF0114(t *testing.T) {
	testToRDF("toRdf-0114-in.jsonld", "toRdf-0114-out.nq", false, "", t)
}

func TestToRDF0115(t *testing.T) {
	testToRDF("toRdf-0115-in.jsonld", "toRdf-0115-out.nq", false, "", t)
}

func TestToRDF0116(t *testing.T) {
	testToRDF("toRdf-0116-in.jsonld", "toRdf-0116-out.nq", false, "", t)
}

func TestToRDF0117(t *testing.T) {
	testToRDF("toRdf-0117-in.jsonld", "toRdf-0117-out.nq", false, "", t)
}

func TestToRDF0118(t *testing.T) {
	testToRDF("toRdf-0118-in.jsonld", "toRdf-0118-out.nq", true,
		"http://json-ld.org/test-suite/tests/", t)
}

func TestToRDF0119(t *testing.T) {
	testToRDF("toRdf-0119-in.jsonld", "toRdf-0119-out.nq", false, "", t)
}
