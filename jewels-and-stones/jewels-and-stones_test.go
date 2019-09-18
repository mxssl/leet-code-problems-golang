package jewelsandstones

import "testing"

func TestCase1(t *testing.T) {
	output := numJewelsInStones1("aA", "aAAbbbb")
	expected := 3
	if output != expected {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	output := numJewelsInStones2("aA", "aAAbbbb")
	expected := 3
	if output != expected {
		t.Fail()
	}
}
