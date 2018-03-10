package main

import "testing"

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
