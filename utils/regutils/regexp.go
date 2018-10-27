package regutils

import (
	"fmt"
	"regexp"
)

func FindSubmatch(s string, regexp *regexp.Regexp) (string, error) {
	match := regexp.FindStringSubmatch(s)
	if len(match) < 2 {
		return "", fmt.Errorf("reg: " + regexp.String() + " missed")
	}
	return match[1], nil
}
