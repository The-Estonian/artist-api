package helpers

import (
	"strings"
)


// Capitalize first letter of the string
func Capitalize(str string) string {
	return strings.ToUpper(string(str[0]))+string(str[1:])
}