package main
import(
	"fmt"
	"math/rand"
	"time"
)
func main() {
	switch tempnum := random(); tempnum {
	case 0, 2, 4, 6, 8:
		fmt.Println("got even number - ", tempnum)
	case 1, 3, 5, 7, 9:
		fmt.Println("got odd number - ", tempnum)
	}
}
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}