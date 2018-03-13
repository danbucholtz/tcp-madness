package main

import (
	"errors"
)

func processCommand(command *Command) (bool, error) {
	if command.Type == IndexNumber {
		return processIndexCommand(command)
	} else if command.Type == RemoveNumber {
		return processRemoveCommand(command)
	} else if command.Type == QueryNumber {
		return processQueryCommand(command)
	}
	return false, errors.New("Unknown Command Type")
}

func processIndexCommand(command *Command) (bool, error) {
	// okay, first thing we need to do is to check if all of the dependencies are indexed
	for i := range command.Dependencies {
		dependency := command.Dependencies[i]
		// if it's not indexed, then return false
		if !IsIndexed(dependency) {
			return false, nil
		}
	}

	IndexPackage(command.PackageName, command.Dependencies)
	// okay cool, we've checked each dependency and they all appear to be indexed, so
	// we're good to go
	return true, nil
}

func processRemoveCommand(command *Command) (bool, error) {
	// okay, first thing we need to do is check if this command is a dependency of another command.
	// if it is, we cannot remove it.
	if IsDependedOn(command.PackageName) {
		// it's depended on, so it can't be removed
		return false, nil
	}

	// w00t, we can remove it
	RemovePackage(command.PackageName)
	return true, nil
}

func processQueryCommand(command *Command) (bool, error) {
	return IsIndexed(command.PackageName), nil
}
