package main

import "fmt"

func merge(left []int, right []int) []int {
	list_sorted := []int{}

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] > right[0] {
				list_sorted = append(list_sorted, right[0])
				list_sorted = append(list_sorted, left[0])
				left = left[1:]
				right = right[1:]
			} else {
				list_sorted = append(list_sorted, left[0])
				list_sorted = append(list_sorted, right[0])
				left = left[1:]
				right = right[1:]
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

func divide(list []int) []int {
	length := len(list)

	if length < 2 {
		return list
	}

	middle := length / 2
	return merge(divide(list[:middle]), divide(list[middle:]))
}

func sort(list []int) []int {
	return divide(list)
}

func main() {
	list := []int{1, 3, 4, 3, 4, 5, 6, 4}
	fmt.Println(sort(list))
}
