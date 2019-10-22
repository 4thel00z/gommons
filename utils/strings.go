package utils

import "strings"

func Any(ss ...string) (result bool) {
	for _, s := range ss {
		result = strings.TrimSpace(s) != "" || result
	}

	return result
}

