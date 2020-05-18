package main

import "fmt"

func pow(pot int, exp int) int {
	if exp == 0 {
		return 1
	}

	if exp < 2 {
		return pot
	}
	return pot * pow(pot, exp-1)
}

func main() {
	var pot, exp int

	fmt.Print("Informe a base e o expoente: ")
	fmt.Scanln(&pot, &exp)
	fmt.Println("Valor:", pow(pot, exp))
}
