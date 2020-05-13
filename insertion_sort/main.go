package main

import "fmt"

func sort(list []int) []int {
	lenght := len(list)
	key, j := 0, 0
	
	if lenght < 2 { return list }

	for i := 1; i < lenght; i++ {
		j = i - 1
		key = list[i]
		
		for j >= 0 && key < list[j] {
			list[j+1] = list[j]
			list[j] = key
			j--
		}
	}

	return list
}


func main() {
	list := []int{21, 5, 4, 6, 3, 2, 1, 9, 10}
	fmt.Println(sort(list))
}