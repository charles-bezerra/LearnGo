package main

import (
	"fmt"
)

func quick_sort(list []int) []int {
	length := len(list)
	end := length - 1

	if length < 2 {
		return list
	}

	pivo_index := 0
	pivo := list[end]

	for i := 0; i < end; i++ {
		if list[i] < pivo {
			list[i], list[pivo_index] = list[pivo_index], list[i]
			pivo_index++
		}
	}

	list[pivo_index], list[end] = list[end], list[pivo_index]

	quick_sort(list[:pivo_index])   //left
	quick_sort(list[pivo_index+1:]) //right

	return list
}

func main() {
	var n int = 0

	fmt.Print("Enter With Numbers Of Entries: ")
	fmt.Scanln(&n)

	list := make([]int, n)

	fmt.Print("List: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}

	fmt.Print("Sorted List:")
	for _, e := range quick_sort(list) {
		fmt.Print(" ", e)
	}

	fmt.Println("")
}
