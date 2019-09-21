package mostcommonword

import (
	"testing"
)

func TestCase1(t *testing.T) {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	expected := "ball"
	output := mostCommonWord(paragraph, banned)
	if output != expected {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	paragraph := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}
	expected := "b"
	output := mostCommonWord(paragraph, banned)
	if output != expected {
		t.Fail()
	}
}
