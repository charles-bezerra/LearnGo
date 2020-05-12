package main

import "fmt"

func sort(list []int, n int) []int {
	
	if n < 1 { return list }

	for i := n-1; i > 0; i-- {
		if list[i] < list[i-1] {
			list[i], list[i-1] = list[i-1], list[i]
		}
	}

	return sort(list, n-1)
}

func main()  {
	list := []int{4, 2, 3, 5, 0, 10}
	fmt.Println(sort(list, len(list)))
}