package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USERNAME")
	fmt.Println("Hello, logged in as", name)
}
