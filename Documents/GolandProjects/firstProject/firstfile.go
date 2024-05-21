package main

import (
	"fmt"
	"sort"
)

func main() {
	var length int
	_, err := fmt.Scanln(&length)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var vector = make([]int, length)
	for i := 0; i < length; i++ {
		var temp int
		_, err := fmt.Scanln(&temp)
		if err != nil {
			fmt.Println("Error reading input:", err)
		}
		vector[i] = temp
	}

	fmt.Println(vector)

	sort.Ints(vector)
	fmt.Println(vector)
}
