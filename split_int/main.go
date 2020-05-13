package main

import "fmt" 

func splitForInt(num uint, list []uint) (uint, []uint) {
	if num < 10 {
		return num, append(list, num)
	}
	return splitForInt(num/10, append(list, num%10))
}


func split(num uint) (list []uint) {
	j := 0
	_, list = splitForInt(num, list)
	for i := len(list)-1; i >= len(list)/2; i-- {
		list[i], list[j] = list[j], list[i]
		j++
	}
	return list
}


func main() {
	fmt.Println( split(123456) )
}
