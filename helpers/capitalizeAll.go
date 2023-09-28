package helpers

import (
	"strings"
)


//Capitalize all first letters of long string
func CapitalizeAll(str string) string {
	strArray := strings.Split(str, " ")
	returnString :=""
	for i := 0; i < len(strArray); i++ {
		returnString += strings.ToUpper(string(string(strArray[i])[0]))+strArray[i][1:] 
		if i < len(strArray)-1{
			returnString += " "
		}
	}
	return returnString
}
