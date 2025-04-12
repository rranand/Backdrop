package validator

import (
	"regexp"
	"strings"
)

func IsEmailValid(txt string) bool {
	txt = TrimString(txt)
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`, txt)
	return match
}

func TrimString(txt string) string {
	return strings.TrimSpace(txt)
}

func IsTextEmpty(txt string) bool {
	return len(TrimString(txt)) == 0
}
