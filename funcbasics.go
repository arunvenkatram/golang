package main
import (
	"fmt"
	"strings"
)
func main() {
	fname := "arun"
	lname := "venkatram"
	fmt.Println(converter(fname, lname))
}
func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)
	return s1, s2
}