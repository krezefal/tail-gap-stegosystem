package tests

import (
	"github.com/krezefal/tail-gap-stegosystem/pkg"
	"testing"
)

func TestPrepareLine(t *testing.T) {
	line1 := "\n"
	line2 := " \n"
	line3 := "actually normal string\n"
	line4 := "actually normal string with space at the end \n"
	line5 := "actually normal string with 2 spaces at the end  \n"

	result := pkg.PrepareLine(line1)
	if result != "" {
		t.Errorf("1. Result was incorrect. Got: |%s|, expected: |%s|", result, "")
	}

	result = pkg.PrepareLine(line2)
	if result != "" {
		t.Errorf("2. Result was incorrect. Got: |%s|, expected: |%s|", result, "")
	}

	result = pkg.PrepareLine(line3)
	if result != "actually normal string" {
		t.Errorf("3. Result was incorrect. Got: |%s|, expected: |%s|", result, "actually normal string")
	}

	result = pkg.PrepareLine(line4)
	if result != "actually normal string with space at the end" {
		t.Errorf("4. Result was incorrect. Got: |%s|, expected: |%s|", result, "actually normal string with space at the end")
	}

	result = pkg.PrepareLine(line5)
	if result != "actually normal string with 2 spaces at the end" {
		t.Errorf("5. Result was incorrect. Got: |%s|, expected: |%s|", result, "actually normal string with 2 spaces at the end")
	}
}
