package termapp

import (
	"testing"
)

func TestTokenizerSpace(t *testing.T) {
	line := "    line "
	tokens := tokenize(line)

	if len(tokens) != 1 {
		t.FailNow()
	}

	var arr0 [1]string
	copy(arr0[:], tokens)
	if [1]string{"line"} != arr0 {
		t.Fail()
	}

	line = "    line arg1     "
	tokens = tokenize(line)

	if len(tokens) != 2 {
		t.Fail()
	}

	var arr1 [2]string
	copy(arr1[:], tokens)
	if [2]string{"line", "arg1"} != arr1 {
		t.Fail()
	}
}
