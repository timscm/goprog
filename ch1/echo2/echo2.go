package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, value := range os.Args {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
}
