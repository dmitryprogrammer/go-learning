package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var i, argsLen int
	argsLen = len(os.Args)
	for i = 1; i < argsLen; i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}