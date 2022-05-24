package entity

import (
	"regexp"
	"strings"
	"unicode"
)

var spacePattern = regexp.MustCompile(`[^\S\t\v\r]`)

type Name string

const maxNameChars = 127

func NewName(name string) (Name, error) {
	if len(name) == 0 {
		return "", NewErrInvalidName(name)
	}

	for i := range name {
		if i+1 > maxNameChars {
			return "", NewErrInvalidName(name)
		}
	}

	tokens := spacePattern.Split(name, -1)

	for _, token := range tokens {
		// empty token signals leading/trailing or an extra space in the input name
		if len(token) == 0 {
			return "", NewErrInvalidName(name)
		}

		for _, chr := range token {
			if !unicode.IsLetter(chr) {
				return "", NewErrInvalidName(name)
			}
		}
	}

	return Name(strings.ToLower(name)), nil
}
