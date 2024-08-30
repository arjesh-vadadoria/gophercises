package main

import (
	"fmt"
)

func main() {
	//printPattern()
	printPatternRec(0, 4)
}

func printPattern() {
	rows := 4
	for r := 0; r < rows; r++ {
		for s := 0; s < rows-(r+1); s++ {
			fmt.Print(" ")
		}
		for c := 0; c < r+1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func space(i int, max int) {
	if i > max {
		fmt.Println()
		return
	}
	if i > 0 {
		fmt.Print(" ")
	}
	space(i+1, max)
}

func printPatternRec(i int, max int) {
	if i > max {
		fmt.Println()
		return
	}
	//space(0, i+1)
	fmt.Print("*")
	printPatternRec(i+1, max)
}
