package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			countilens(f, counts)
			f.Close()
		}
	}

	var uniqueCounts []int
	var line int = 0

	for _, n := range counts {
		if n != line {
			line = n
			uniqueCounts = append(uniqueCounts, line)
		}
	}

	for _, n := range uniqueCounts {
		fmt.Printf("%d\n", n)
	}
}

func countilens(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
