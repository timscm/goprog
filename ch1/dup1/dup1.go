package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Linux: Ctrl+D, stop input
	// Windows: Ctrl+Z, stop input
	fmt.Println("Ready to input:")
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println("input end, start to show")
	// only show dup lines.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
