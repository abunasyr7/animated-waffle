package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return false
	}

	for _, token := range strings.Fields(s) {
		if token == "" {
			continue
		}

		for _, r := range token {
			if r != '.' && r != '-' {
				return false
			} 
		}
	}

	return true
}

func Converter (s string) string {
	if isMorse(s) {
		return morse.ToText(s)
	}

	return morse.ToMorse(s)
}