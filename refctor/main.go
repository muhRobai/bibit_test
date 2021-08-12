package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(findFirstStringInBracket("(testsasasas)"))
}

func findFirstStringInBracket(str string) string {
	if len(str) < 0 {
		return ""
	}

	var strResult string
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound >= 0 {
		index := strings.Index(str[indexFirstBracketFound:], ")")
		if index <= 0 {
			return ""
		}
		strResult = str[1:index]
	}

	return strResult
}
