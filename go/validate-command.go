package main

import (
	"errors"
	"strings"
)

// Command is a struct built from the raw string
type Command struct {
	Type         int
	PackageName  string
	Dependencies []string
}

// ValidateCommand Returns a command if valid, otherwise returns an error
func ValidateCommand(rawCommand string) (Command, error) {
	var command Command

	if len(rawCommand) == 0 {
		return command, errors.New("Usage: Invalid Command, command must be a string")
	}

	lastIndex := len(rawCommand) - 1
	lastCharacter := string(rawCommand[lastIndex])

	if lastCharacter != "\n" {
		return command, errors.New("Usage: Commands must end in a newline character")
	}

	commandChunks := strings.Split(rawCommand, "|")
	if len(commandChunks) != 3 {
		return command, errors.New("Usage: command|package|dep1,dep2,dep3,... ")
	}

	var cleanedChunks []string
	for i := range commandChunks {
		trimmed := strings.Trim(commandChunks[i], " ")
		if len(trimmed) != 0 {
			cleanedChunks = append(cleanedChunks, trimmed)
		}
	}

	commandType := cleanedChunks[0]
	if commandType != "INDEX" && commandType != "REMOVE" && commandType != "QUERY" {
		return command, errors.New("Usage: First argument must be INDEX, REMOVE, or QUERY")
	}

	return command, errors.New("TODO")

}
