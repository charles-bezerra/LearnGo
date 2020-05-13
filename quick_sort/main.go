package main


import (
	"fmt"
	"os"
	"/home/charles/Desktop/LearnGo/quick_sort/util/util"
)

func quick(left []int, right []int) []int {
	return left
}

func sort(list []int) []int {
	return list
}

func main()  {
	path, error := os.Getwd()

	
	if error != nil  {
		fmt.Println(error)
	}

	fmt.Println(path)
}