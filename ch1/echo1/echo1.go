package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i:=1; i<len(os.Args); i++ {
		s+= sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// for loop
	var i = 0
	for {
		fmt.Println("will loop 3 times")
		if i >= 3 {
			break
		}
		i++
	}
}
