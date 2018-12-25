package main

import (
	"fmt"
)

func main() {
	arunmap := make(map[string]int)
	arunmap["go"] = 2
	arunmap["python"] = 5
	venkatmap := map[string]int{
		"docker":     6,
		"Kubernetes": 9,
	}

	fmt.Printf("Arun is %v\nVenkat is %v", arunmap, venkatmap)
}
