package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("Hello, my name is " + os.Args[1])
}