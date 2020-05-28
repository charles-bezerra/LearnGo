package main

import (
	"fmt"
)

func merge(left []int, right []int) []int {
	list_sorted := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] > right[0] {
				list_sorted = append(list_sorted, right[0])
				right = right[1:]

			} else {
				list_sorted = append(list_sorted, left[0])
				left = left[1:]

			}
		}

		if len(left) > 0 && len(right) == 0 {
			list_sorted = append(list_sorted, left[0])
			left = left[1:]
		}

		if len(right) > 0 && len(left) == 0 {
			list_sorted = append(list_sorted, right[0])
			right = right[1:]
		}
	}
	return list_sorted
}

func mergeSort(list []int) []int {
	length := len(list)
	if length < 2 {
		return list
	}
	middle := length / 2
	return merge(mergeSort(list[:middle]), mergeSort(list[middle:]))
}

func sort(list []int) []int {
	return mergeSort(list)
}

func main() {
	n := 0
	fmt.Print("Enter with number of inputs: ")
	fmt.Scanln(&n)

	list := make([]int, n)
	fmt.Print("List: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}

	fmt.Println("\nSorted list:", sort(list))
}
