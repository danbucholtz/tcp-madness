package main

import (
	"errors"
)

func processCommand(command *Command) (bool, error) {
	if command.Type == INDEX_NUMBER {
		return processIndexCommand(command)
	} else if command.Type == REMOVE_NUMBER {
		return processRemoveCommand(command)
	} else if command.Type == QUERY_NUMBER {
		return processQueryCommand(command)
	}
	return false, errors.New("Unknown Command Type")
}

func processIndexCommand(command *Command) (bool, error) {
	Debugf("Starting to process index command for %s", command.PackageName)
	// okay, first thing we need to do is to check if all of the dependencies are indexed
	for i := range command.Dependencies {
		dependency := command.Dependencies[i]
		// if it's not indexed, then return false
		if !IsIndexed(dependency) {
			Debugf("Failed to process index command for %s due to dependency %s not being indexed", command.PackageName, dependency)
			return false, nil
		}
	}

	IndexPackage(command.PackageName, command.Dependencies)
	// okay cool, we've checked each dependency and they all appear to be indexed, so
	// we're good to go
	Debugf("Successfully processed index command for %s", command.PackageName)
	return true, nil
}

func processRemoveCommand(command *Command) (bool, error) {
	Debugf("Starting to process remove command for %s", command.PackageName)
	// okay, first thing we need to do is check if this command is a dependency of another command.
	// if it is, we cannot remove it.
	if IsDependedOn(command.PackageName) {
		// it's depended on, so it can't be removed
		Debugf("Failed to process remove command for %s since it's depended on by another package", command.PackageName)
		return false, nil
	}

	// w00t, we can remove it
	RemovePackage(command.PackageName)
	Debugf("Successfully processed remove command for %s", command.PackageName)
	return true, nil
}

func processQueryCommand(command *Command) (bool, error) {
	Debugf("Starting to process query command for %s", command.PackageName)
	result := IsIndexed(command.PackageName)
	Debugf("Successfully processed query command for %s, result is %t", command.PackageName, result)
	return result, nil
}
