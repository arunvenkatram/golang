package main
import(
	"fmt"
)
func main() {
	switch "docker" {
	case "linux":
		fmt.Println("this is case block of linux")
	case "docker":
		fmt.Println("This is case block of docker")
	case "windows":
		fmt.Println("this is case block of windows")
	}
}