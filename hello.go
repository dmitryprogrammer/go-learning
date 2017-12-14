package main

import (
	"os"
	"fmt"
)

func main() {
	a3 := [20]int{1,2,3,4}
	fmt.Println("Hello, my name is " + os.Args[1] + "\n")
	fmt.Println(a3)
}

