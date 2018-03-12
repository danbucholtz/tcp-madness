package main

import (
	"regexp"
	"strings"
)

// ReduceBools takes a list of booleans and ANDs them
func ReduceBools(bools []bool) bool {
	for i := 0; i < len(bools); i++ {
		if bools[i] == false {
			return false
		}
	}
	return true
}

// IsValidUnixCommand determines whether a string is a valid-ish unix command name. Note, this does not actually validate that it's an executable command, or installed on ones system, etc
func IsValidUnixCommand(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_+-]+$", s)
	return match
}

// TrimAndFilter takes a list of strings, runs a trim() operation on them, then filters out any empty strings
func TrimAndFilter(list []string) []string {
	var toReturn []string
	for i := range list {
		noNewLine := strings.Trim(list[i], "\n")
		trimmed := strings.Trim(noNewLine, " ")
		if len(trimmed) != 0 {
			toReturn = append(toReturn, trimmed)
		}
	}
	return toReturn
}
