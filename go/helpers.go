package main

import (
	"regexp"
)

func IsValidUnixCommand(s string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z0-9_+-]+$", s)
	return match
}
