package main

import (
	"fmt"
)

func main() {
	printHollowBox()
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

func printPatternK() {
	rows := 4
	for r := 0; r < rows; r++ {
		for c := 0; c < rows-(r); c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < r+1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printFilledBox() {
	rows := 4
	for r := 0; r < rows; r++ {
		for c := 0; c < rows; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printHollowBox() {
	rows := 8
	for r := 0; r < rows; r++ {
		for c := 0; c < rows; c++ {
			if r == 0 || r == rows-1 || c == 0 || c == rows-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
			//if r%2 == 0 {
			fmt.Print(" ")
			//}
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

func matrix() {
	list1 := []string{"1", "2", "3"}
	list2 := []string{"4", "5", "6"}
	list3 := []string{"A", "B", "C"}
	var result [][]string

	for i := 0; i < 3; i++ {
		var tempList []string
		tempList = append(tempList, list1[0])
		tempList = append(tempList, list2[0])
		tempList = append(tempList, list3[0])
		list1 = list1[1:]
		list2 = list2[1:]
		list3 = list3[1:]
		result = append(result, tempList)
		fmt.Println(tempList)
	}
	fmt.Println(result)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
