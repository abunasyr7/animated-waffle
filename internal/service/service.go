package service

import (
	"strings"
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
