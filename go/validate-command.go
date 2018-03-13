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
func ValidateCommand(rawCommand string) (*Command, error) {
	var command *Command

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
		return command, errors.New("Usage: command|package|dep1,dep2,dep3,...")
	}

	cleanedChunks := TrimAndFilter(commandChunks)

	commandType := cleanedChunks[0]
	if commandType != "INDEX" && commandType != "REMOVE" && commandType != "QUERY" {
		return command, errors.New("Usage: First argument must be INDEX, REMOVE, or QUERY")
	}

	if !IsValidUnixCommand(cleanedChunks[1]) {
		return command, errors.New("Usage: Commands must be a valid unix command format")
	}

	dependencies := []string{}

	if len(cleanedChunks) > 2 {
		rawDependencies := strings.Split(cleanedChunks[2], ",")
		dependencies = TrimAndFilter(rawDependencies)

		if len(rawDependencies) != len(dependencies) {
			return command, errors.New("Usage: One or more invalid dependency entries")
		}
	}

	command = &Command{
		Type:         commandToNumber(commandType),
		PackageName:  cleanedChunks[1],
		Dependencies: dependencies,
	}

	return command, nil
}

func commandToNumber(command string) int {
	if command == Index {
		return IndexNumber
	} else if command == Remove {
		return RemoveNumber
	} else if command == Query {
		return QueryNumber
	}
	return UnknownNumber
}
