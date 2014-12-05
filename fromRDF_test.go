package gojsonld

import (
	"encoding/json"
	"reflect"
	"testing"
)

func testFromRDF(input, output string, useNativeTypes, useRdfType bool, t *testing.T) {
	inputRDF, inputErr := ReadDatasetFromFile(test_dir + input)
	if !isNil(inputErr) {
		t.Error("Could not open input: " + inputErr.Error())
		return
	}
	outputJSON, outputErr := ReadJSONFromFile(test_dir + output)
	if !isNil(outputErr) {
		t.Error("Could not open output: " + outputErr.Error())
		return
	}
	options := &Options{}
	options.UseNativeTypes = useNativeTypes
	options.UseRdfType = useRdfType
	serializedJSON := FromRDF(inputRDF, options)
	serializedString, _ := json.MarshalIndent(serializedJSON, "", "    ")
	outputString, _ := json.MarshalIndent(outputJSON, "", "    ")
	if !reflect.DeepEqual(serializedJSON, outputJSON) {
		t.Error("Expected:\n", string(outputString), "\nGot:\n",
			string(serializedString))
		return
	}
}

func TestFromRDF0001(t *testing.T) {
	testFromRDF("fromRdf-0001-in.nq", "fromRdf-0001-out.jsonld", false, false, t)
}

func TestFromRDF0002(t *testing.T) {
	testFromRDF("fromRdf-0002-in.nq", "fromRdf-0002-out.jsonld", false, false, t)
}

func TestFromRDF0003(t *testing.T) {
	testFromRDF("fromRdf-0003-in.nq", "fromRdf-0003-out.jsonld", false, false, t)
}

func TestFromRDF0004(t *testing.T) {
	testFromRDF("fromRdf-0004-in.nq", "fromRdf-0004-out.jsonld", false, false, t)
}

func TestFromRDF0005(t *testing.T) {
	testFromRDF("fromRdf-0005-in.nq", "fromRdf-0005-out.jsonld", false, false, t)
}

func TestFromRDF0006(t *testing.T) {
	testFromRDF("fromRdf-0006-in.nq", "fromRdf-0006-out.jsonld", false, false, t)
}

func TestFromRDF0007(t *testing.T) {
	testFromRDF("fromRdf-0007-in.nq", "fromRdf-0007-out.jsonld", false, false, t)
}

func TestFromRDF0008(t *testing.T) {
	testFromRDF("fromRdf-0008-in.nq", "fromRdf-0008-out.jsonld", false, false, t)
}

func TestFromRDF0009(t *testing.T) {
	testFromRDF("fromRdf-0009-in.nq", "fromRdf-0009-out.jsonld", false, false, t)
}

func TestFromRDF0010(t *testing.T) {
	testFromRDF("fromRdf-0010-in.nq", "fromRdf-0010-out.jsonld", false, false, t)
}

func TestFromRDF0011(t *testing.T) {
	testFromRDF("fromRdf-0011-in.nq", "fromRdf-0011-out.jsonld", false, false, t)
}

func TestFromRDF0012(t *testing.T) {
	testFromRDF("fromRdf-0012-in.nq", "fromRdf-0012-out.jsonld", false, false, t)
}

func TestFromRDF0013(t *testing.T) {
	testFromRDF("fromRdf-0013-in.nq", "fromRdf-0013-out.jsonld", false, false, t)
}

func TestFromRDF0014(t *testing.T) {
	testFromRDF("fromRdf-0014-in.nq", "fromRdf-0014-out.jsonld", false, false, t)
}

func TestFromRDF0015(t *testing.T) {
	testFromRDF("fromRdf-0015-in.nq", "fromRdf-0015-out.jsonld", false, false, t)
}

func TestFromRDF0016(t *testing.T) {
	testFromRDF("fromRdf-0016-in.nq", "fromRdf-0016-out.jsonld", false, false, t)
}

func TestFromRDF0017(t *testing.T) {
	testFromRDF("fromRdf-0017-in.nq", "fromRdf-0017-out.jsonld", false, false, t)
}

func TestFromRDF0018(t *testing.T) {
	testFromRDF("fromRdf-0018-in.nq", "fromRdf-0018-out.jsonld", true, false, t)
}

func TestFromRDF0019(t *testing.T) {
	testFromRDF("fromRdf-0019-in.nq", "fromRdf-0019-out.jsonld", false, true, t)
}

func TestFromRDF0020(t *testing.T) {
	testFromRDF("fromRdf-0020-in.nq", "fromRdf-0020-out.jsonld", false, false, t)
}

func TestFromRDF0021(t *testing.T) {
	testFromRDF("fromRdf-0021-in.nq", "fromRdf-0021-out.jsonld", false, false, t)
}

func TestFromRDF0022(t *testing.T) {
	testFromRDF("fromRdf-0022-in.nq", "fromRdf-0022-out.jsonld", false, false, t)
}
