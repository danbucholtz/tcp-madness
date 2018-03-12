package main

import (
	"testing"
)

func TestValidateCommand_fullCommandNoEndline(t *testing.T) {
	expectedErrorMsg := "Usage: Commands must end in a newline character"
	_, err := ValidateCommand("taco|time|yo")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_tacoOnly(t *testing.T) {
	expectedErrorMsg := "Usage: Commands must end in a newline character"
	_, err := ValidateCommand("taco")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_onlyNewline(t *testing.T) {
	expectedErrorMsg := "Usage: command|package|dep1,dep2,dep3,..."
	_, err := ValidateCommand("\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_tacoThenNewLine(t *testing.T) {
	expectedErrorMsg := "Usage: command|package|dep1,dep2,dep3,..."
	_, err := ValidateCommand("taco\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_tacoPipeNewline(t *testing.T) {
	expectedErrorMsg := "Usage: command|package|dep1,dep2,dep3,..."
	_, err := ValidateCommand("taco|\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_tacoPipeTimeNewline(t *testing.T) {
	expectedErrorMsg := "Usage: command|package|dep1,dep2,dep3,..."
	_, err := ValidateCommand("taco|time\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_tacoPipeTimePipeForPipeEveryoneNewline(t *testing.T) {
	expectedErrorMsg := "Usage: command|package|dep1,dep2,dep3,..."
	_, err := ValidateCommand("taco|time|for|everyone\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_invalidFirstArg(t *testing.T) {
	expectedErrorMsg := "Usage: First argument must be INDEX, REMOVE, or QUERY"
	_, err := ValidateCommand("taco|time|\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func TestValidateCommand_basicCommand(t *testing.T) {
	command, _ := ValidateCommand("INDEX|ls|\n")
	if command.PackageName != "ls" {
		t.Error("Expected package name to be ls")
	}
	if len(command.Dependencies) != 0 {
		t.Error("Expected empty array for dependencies")
	}
	if command.Type != IndexNumber {
		t.Errorf("Expected %d for the type", IndexNumber)
	}
}

func TestValidateCommand_basicWithDeps(t *testing.T) {
	command, _ := ValidateCommand("INDEX|ls|cd,cat,man,tsc,node,npm\n")
	//fmt.Println(command)
	if command.PackageName != "ls" {
		t.Error("Expected package name to be ls")
	}
	if len(command.Dependencies) != 6 {
		t.Error("Expected six dependencies")
	}
	validateDeps(t, command.Dependencies, 0, "cd")
	validateDeps(t, command.Dependencies, 1, "cat")
	validateDeps(t, command.Dependencies, 2, "man")
	validateDeps(t, command.Dependencies, 3, "tsc")
	validateDeps(t, command.Dependencies, 4, "node")
	validateDeps(t, command.Dependencies, 5, "npm")
	if command.Type != IndexNumber {
		t.Errorf("Expected %d for the type", IndexNumber)
	}
}

func TestValidateCommand_strangeDeps(t *testing.T) {
	command, _ := ValidateCommand("INDEX|ls|cd, cat, man,tsc, node,npm\n")
	if command.PackageName != "ls" {
		t.Error("Expected package name to be ls")
	}
	if len(command.Dependencies) != 6 {
		t.Error("Expected six dependencies")
	}
	validateDeps(t, command.Dependencies, 0, "cd")
	validateDeps(t, command.Dependencies, 1, "cat")
	validateDeps(t, command.Dependencies, 2, "man")
	validateDeps(t, command.Dependencies, 3, "tsc")
	validateDeps(t, command.Dependencies, 4, "node")
	validateDeps(t, command.Dependencies, 5, "npm")
	if command.Type != IndexNumber {
		t.Errorf("Expected %d for the type", IndexNumber)
	}
}

func TestValidateCommand_emptyCsvEntryInDeps(t *testing.T) {
	expectedErrorMsg := "Usage: One or more invalid dependency entries"
	_, err := ValidateCommand("INDEX|ls|cd,,man, tsc,,node, npm\n")
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected to receive %s, received %s", expectedErrorMsg, err.Error())
	}
}

func validateDeps(t *testing.T, dependencies []string, index int, value string) {
	if dependencies[index] != value {
		t.Errorf("Expected dependencies[%d] to equal %s\n", index, value)
	}
}
