package utils

import "strings"

func Repeat(character string) string {
	var ans strings.Builder
	for i:=0;i<1000;i++ {
		ans.WriteString(character)
	}
	return  ans.String()
}