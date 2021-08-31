// go run dup2.go 1.txt 2.txt 3.txt
// or
// go run dup2.go
// > line 1
// > line 1
// > line 2
// > line 2
// > line 3
// > Ctrl+D
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	fmt.Println("Start to output:")
	// deal output
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("dup line: %d\t%s\n", n, line)
		} else {
			fmt.Printf("sig line: %d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	fmt.Println("Ready for input content:")
	for input.Scan() {
		counts[input.Text()]++
	}
}
