package main

import "fmt"

func merge(left []int, right []int) []int {
	list := []int{}
	
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if right[0] < left[0] {
				list = append(list, right[0])
				right = right[1:]
			} else {
				list = append(list, left[0])
				left = left[1:]			
			}
		}
		
		if len(right) == 0 && len(left) > 0 {
			list = append(list, left[0])
			left = left[1:]
		}

	 	if len(left) == 0 && len(right) > 0 {
			list = append(list, right[0])
			right = right[1:]
		}
	}

	return list
}
 
func sort(list []int) []int {
	length := len(list)
	
	if length < 2 { 
		return list 
	}

	middle := length/2 
	return merge(sort(list[:middle]), sort(list[middle:]))
}

func main() {
	list := []int{1,5,6,10,2,50,12, 39, 12, 12}
	fmt.Println(sort(list))
}