package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var i, argsLen int
	argsLen = len(os.Args)

	if os.Args[1] == "a" {
		s = ""
		sep = ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	} else {
		for i = 1; i < argsLen; i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}
}