package main

import (
	"testing"
)

func TestIsValidUnixCommandInvalid(t *testing.T) {
	resultOne := IsValidUnixCommand("abc ")
	if resultOne {
		t.Error("'abc ' is not a valid unix command but is registering as one")
	}

	resultTwo := IsValidUnixCommand("abc?")
	if resultTwo {
		t.Error("'abc?' is not a valid unix command but is registering as one")
	}

	resultThree := IsValidUnixCommand("")
	if resultThree {
		t.Error("'' is not a valid unix command but is registering as one")
	}

	resultFour := IsValidUnixCommand(" ")
	if resultFour {
		t.Error("' ' is not a valid unix command but is registering as one")
	}

	resultFive := IsValidUnixCommand("(")
	if resultFive {
		t.Error("'(' is not a valid unix command but is registering as one")
	}

	resultSix := IsValidUnixCommand("ABC ")
	if resultSix {
		t.Error("'ABC ' is not a valid unix command but is registering as one")
	}

	resultSeven := IsValidUnixCommand("emacs elisp")
	if resultSeven {
		t.Error("'emacs elisp' is not a valid unix command but is registering as one")
	}
}

func TestIsValidUnixCommandValid(t *testing.T) {
	resultOne := IsValidUnixCommand("abc")
	if !resultOne {
		t.Error("'abc' is a valid unix command but is not registering as one")
	}

	resultTwo := IsValidUnixCommand("q")
	if !resultTwo {
		t.Error("'q' is a valid unix command but is not registering as one")
	}

	resultThree := IsValidUnixCommand("ABC")
	if !resultThree {
		t.Error("'ABC' is a valid unix command but is not registering as one")
	}

	resultFour := IsValidUnixCommand("Q")
	if !resultFour {
		t.Error("'Q' is a valid unix command but is not registering as one")
	}

	resultFive := IsValidUnixCommand("abc123")
	if !resultFive {
		t.Error("'abc123' is a valid unix command but is not registering as one")
	}

	resultSix := IsValidUnixCommand("abc_123")
	if !resultSix {
		t.Error("'abc_123' is a valid unix command but is not registering as one")
	}

	resultSeven := IsValidUnixCommand("abc-123")
	if !resultSeven {
		t.Error("'abc-123' is a valid unix command but is not registering as one")
	}

	resultEight := IsValidUnixCommand("abc+123")
	if !resultEight {
		t.Error("'abc+123' is a valid unix command but is not registering as one")
	}
}

func TestReduceBoolsFalse(t *testing.T) {

	firstTest := []bool{false, false, false}
	resultOne := ReduceBools(firstTest)
	if resultOne {
		t.Error("Failed to reduce booleans from list of false to a single false")
	}

	secondTest := []bool{true, true, false, true}
	resultTwo := ReduceBools(secondTest)
	if resultTwo {
		t.Error("Failed to reduce booleans from mixed list w/ a false to a single false")
	}

	thirdTest := []bool{false}
	resultThree := ReduceBools(thirdTest)
	if resultThree {
		t.Error("Failed to reduce single boolean to a single boolean")
	}
}

func TestReduceBoolesTrue(t *testing.T) {
	firstTest := []bool{true}
	resultOne := ReduceBools(firstTest)
	if !resultOne {
		t.Error("Failed to reduce booleans from single true to single true")
	}

	secondTest := []bool{true, true, true}
	resultTwo := ReduceBools(secondTest)
	if !resultTwo {
		t.Error("Failed to reduce booleans from multiple true to single true")
	}
}

func TestTrimAndFilter_noChange(t *testing.T) {
	firstTest := []string{"first", "second"}
	resultOne := TrimAndFilter(firstTest)
	if len(resultOne) != 2 {
		t.Error("Expected a length of two")
	}
	if resultOne[0] != "first" {
		t.Error("Unexpected first input")
	}

	if resultOne[1] != "second" {
		t.Error("Unexpected second input")
	}
}

func TestTrimAndFilter_trim(t *testing.T) {
	firstTest := []string{"first ", " second \n"}
	resultOne := TrimAndFilter(firstTest)
	if len(resultOne) != 2 {
		t.Error("Expected a length of two")
	}
	if resultOne[0] != "first" {
		t.Error("Unexpected first input")
	}

	if resultOne[1] != "second" {
		t.Error("Unexpected second input")
	}
}
