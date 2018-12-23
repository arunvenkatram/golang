package main
import (
	"fmt"
)
func main() {
	lowestnum := findlowest(13, 24, 3, 76, 45, 4, 76)
	fmt.Println(lowestnum)
}
func findlowest(numbers ...int) (lowestnum int){
	lower := numbers[0]
	for _, i := range numbers {
		if i < lower {
			lower = i
		}
	}
	return lower
}