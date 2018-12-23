package main
import (
	"fmt"
)
func main() {
	arun := "5"
	venkat := "5"
	if arun > venkat {
		fmt.Println("arun is big")
	}else if venkat > arun {
		fmt.Println("venkat is big")
	}else{
		fmt.Println("Both are equal")
	}
}