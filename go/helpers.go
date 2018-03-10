package main

import (
	"regexp"
)

// IsValidUnixCommand determines whether a string is a valid-ish unix command name. Note, this does not actually validate that it's an executable command, or installed on ones system, etc
func IsValidUnixCommand(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_+-]+$", s)
	return match
}
