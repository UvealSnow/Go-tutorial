package greetings

import (
	"regexp"
	"slices"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Errorf(`Hello("%v") = %q, %v want "", error`, name, msg, err)
	}
}

func TestParseNamesList(t *testing.T) {
	names := "Lenina Crown, Bernard Marx"
	want := []string{"Lenina Crown", "Bernard Marx"}
	parsedNames, err := ParseNames(names)

	if !slices.Equal(parsedNames, want) || err != nil {
		t.Errorf("ParseNames(%q) = %q, want %q, nil", names, parsedNames, want)
	}
}

func TestParseNamesEmpty(t *testing.T) {
	names := ""
	parsedNames, err := ParseNames(names)

	if len(parsedNames) > 0 || err == nil {
		t.Errorf(`ParseNames(%q) = %q, %v want "", error`, names, parsedNames, err)
	}
}
