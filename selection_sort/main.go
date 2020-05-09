package main

import "fmt"

func sort (list []int) []int {
	min := 0
	length := len(list)

	if length < 2 { 
		return list 
	}

	for i := 0; i < length-1; i++ {
		min = i

		for j := i+1; j < length; j++ {
			if list[j] < list[min] {
				min = j
			}
		}

		list[i], list[min] = list[min], list[i]
	}

	return list
}

func main()  {
	list := []int{2,39,4,1,5,6}	
	fmt.Println(sort(list))
}