package main

import (
	"fmt"
)

func main() {
	arun := []string{"linux", "docker", "kubernetes", "python", "golang"}
	fmt.Println(arun)
	SliceOfSlice := arun[1:3]
	fmt.Println(SliceOfSlice)
}
