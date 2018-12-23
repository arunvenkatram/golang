package main
import(
	"fmt"
)
func main() {
	fruits := []string{"apple", "oranges", "mangoes"}
	for _, i := range fruits {
		fmt.Println(i)
	}

	for index, i := range fruits {
		fmt.Println("Index is", index, "Value is", i)
	}
}