package tests

import (
	"github.com/krezefal/tail-gap-stegosystem/pkg"
	"testing"
)

func TestPrepareLine(t *testing.T) {
	line1 := "\n"
	line2 := " \n"
	line3 := "\r\n"
	line4 := " \r\n"
	line5 := "actually normal string\n"
	line6 := "actually normal string with space at the end \n"
	line7 := "actually normal string with carriage return\r\n"
	line8 := "actually normal string with space at the end and carriage return \r\n"

	result, crFlag := pkg.PrepareLine(line1)
	if result != "" && crFlag != false {
		t.Errorf("1. Result is incorrect. Got: |%s|, expected: |%s|", result, "")
	}

	result, crFlag = pkg.PrepareLine(line2)
	if result != "" && crFlag != false {
		t.Errorf("1. Result is incorrect. Got: |%s|, expected: |%s|", result, "")
	}

	result, crFlag = pkg.PrepareLine(line3)
	if result != "" && crFlag != true {
		t.Errorf("3. Result is incorrect. Got: |%s|, expected: |%s|", result, "")
	}

	result, crFlag = pkg.PrepareLine(line4)
	if result != "actually normal string" && crFlag != true {
		t.Errorf("4. Result is incorrect. Got: |%s|, expected: |%s|", result, "")
	}

	result, crFlag = pkg.PrepareLine(line5)
	if result != "actually normal string" && crFlag != false {
		t.Errorf("5. Result is incorrect. Got: |%s|, expected: |%s|", result, "actually normal string")
	}

	result, crFlag = pkg.PrepareLine(line6)
	if result != "actually normal string with space at the end" && crFlag != false {
		t.Errorf("6. Result is incorrect. Got: |%s|, expected: |%s|", result, "actually normal string with space at the end")
	}

	result, crFlag = pkg.PrepareLine(line7)
	if result != "actually normal string with carriage return" && crFlag != true {
		t.Errorf("7. Result is incorrect. Got: |%s|, expected: |%s|", result, "actually normal string with carriage return")
	}

	result, crFlag = pkg.PrepareLine(line8)
	if result != "actually normal string with space at the end and carriage return" && crFlag != true {
		t.Errorf("8. Result is incorrect. Got: |%s|, expected: |%s|", result, "actually normal string with space at the end and carriage return")
	}
}
