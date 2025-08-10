package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rangeIteration()
	indexedIteration(true)
}

// First implementation of go command arguments output iteration
func simpleIteration() {
	var arguments, sep string
	for i := 1; i < len(os.Args); i++ {
		arguments += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(arguments)
}

// Iteration implemented with range
func rangeIteration() {
	var arguments string
	var sep = " "

	for _, arg := range os.Args[1:] {
		arguments += arg + sep
	}
	fmt.Println(arguments)
}

// It is not real iteration, but short case how it possible to output arguemnts	
func joinIteration() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Simpliest case how to output arugment
func shortIteration() {
	fmt.Println(os.Args[1:])
}

// Command arguments output with filename, or without depend on argument flag
func indexedIteration(isFileVisible bool) {
	var arguemnts string
	var sep = " "
	var startRange int8

	if isFileVisible {
		startRange = 0
	} else {
		startRange = 1
	}

	for _, arg:= range os.Args[startRange:] {
		arguemnts += arg + sep
	}
	fmt.Println(arguemnts)
}
