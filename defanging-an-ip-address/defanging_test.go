package defanging

import "testing"

func TestIPValidation1(t *testing.T) {
	output := isIPValid("1.1.1.1")
	expected := true
	if output != expected {
		t.Fail()
	}
}

func TestIPValidation2(t *testing.T) {
	output := isIPValid("1.1")
	expected := false
	if output != expected {
		t.Fail()
	}
}

func TestDefanging(t *testing.T) {
	output := defangIPaddr("1.1.1.1")
	expected := "1[.]1[.]1[.]1"
	if output != expected {
		t.Fail()
	}
}

func TestDefangingStrings(t *testing.T) {
	output := defangIPaddrStrings("1.1.1.1")
	expected := "1[.]1[.]1[.]1"
	if output != expected {
		t.Fail()
	}
}
