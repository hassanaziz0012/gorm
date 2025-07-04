package builder

import (
	"regexp"
	"strings"
)

func removeExtraSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}
