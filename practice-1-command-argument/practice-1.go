package main

import (
	"fmt"
	"os"
)

func main() {
	var arguments, sep string
	for i := 1; i < len(os.Args); i++ {
		arguments += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(arguments)
}
