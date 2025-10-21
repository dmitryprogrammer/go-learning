package main

import (
	"fmt"
	"os"
)

func main() {
	a3 := [20]int{1, 2, 3, 4}
	if len(os.Args) > 1 {
		fmt.Printf(fmt.Sprintf("Hello, my name is %s \n", os.Args[1]))
	}
	fmt.Println(a3)
}
