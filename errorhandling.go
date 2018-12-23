package main
import(
	"fmt"
	"os"
)
func main() {
	_, err := os.Open("C:\\arun\\test.txt.txt")
	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}