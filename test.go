package main

import "fmt"

func main()  {
	const v int = 120
	fmt.Printf("%T - %v\n", v, v)

	list := []int{1}

	fmt.Println(list[1:])
}