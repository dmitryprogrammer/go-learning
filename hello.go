package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[:1]) == 0 {
		return
	}

	var userHeight float64
	var userWidth float64
	var errHeight error
	var errWidth error

	userHeight, errHeight = strconv.ParseFloat(os.Args[1], 64)
	userWidth, errWidth = strconv.ParseFloat(os.Args[2], 64)

	if errHeight != nil || errWidth != nil {
		fmt.Print(errHeight, errWidth)
		return
	}

	var IMT = userWidth / math.Pow(userHeight, 2)
	var IMTFormatted = fmt.Sprintf("%.2f", IMT)

	fmt.Print(IMTFormatted)
}
