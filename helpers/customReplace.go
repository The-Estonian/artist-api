package helpers

import (
	"strings"
)

//Replace letters with letters in string input (-1 on all ocasions)
func Replace(str string, what string, withWhat string) string {
	return strings.Replace(str, what, withWhat, -1)
}