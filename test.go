package main

import "fmt"

func change(list []int) {
	list[0] = 0
}

func main() {
	// const v int = 120
	// fmt.Printf("%T - %v\n", v, v)

	// j := []int{1, 2, 3, 4, 5}

	// for len(j) > 0 {
	// 	fmt.Println(j)
	// 	j = j[1:]
	// }

	l := []int{1, 2, 3, 4}
	change(l)
	fmt.Println(l[:1])
}
