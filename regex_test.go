package gojsonld

import (
	"testing"
)

func TestRegexStatement01(t *testing.T) {
	matches := STATEMENT.MatchString("_:alice <http://xmlns.com/foaf/0.1/knows> _:bob <http://example.org/graphs/john> .")
	if !matches {
		t.Error("Input does not match regex")
	}
}

func TestRegexStatement02(t *testing.T) {
	matches := STATEMENT.MatchString("_:alice <http://xmlns.com/foaf/0.1/knows> _:bob .")
	if !matches {
		t.Error("Input does not match regex")
	}
}

func TestRegexStatement03(t *testing.T) {
	matches := STATEMENT.MatchString("_:alice <http://example.org/graphs/john> .")
	if matches {
		t.Error("Input matches regex")
	}
}

func TestRegexIRI01(t *testing.T) {
	matches := IRIREF.MatchString("<http://google.com/>")
	if !matches {
		t.Error("Input does not match regex")
	}
}

func TestRegexIRI02(t *testing.T) {
	matches := IRIREF.MatchString("<http://\"google.com/>")
	if matches {
		t.Error("Input matches regex")
	}
}
