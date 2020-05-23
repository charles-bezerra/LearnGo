package main

import (
	"fmt"
)

func quick_sort(list []int) []int {
	length := len(list)

	if length < 2 {
		return list
	}

	pivo_index := length - 1
	pivo := list[pivo_index]
	i := length - 2

	for i >= 0 && pivo < list[i] {
		list[i], list[i+1] = pivo, list[i]
		pivo_index = i
		i--
	}

	if length > 2 {
		left := quick_sort(list[:pivo_index])
		right := quick_sort(list[pivo_index+1:])

		j := 0

		for i = 0; i < length; i++ {
			if i < pivo_index {
				list[i] = left[i]

			} else if i > pivo_index {
				list[i] = right[j]
				j++

			}
		}
	}

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
