package main

import "fmt"

func soma(number1 int, number2 int) int {
	var result int
	result = number1 + number2
   	return result
}

func main() {
	fmt.Print("Digite o primeiro número: ")
	var number1 int
	fmt.Scanf("%d", &number1)
	
	fmt.Println("Digite o segundo número: ")
	var number2 int
	fmt.Scanf("%d", &number2)

	fmt.Println(soma(number1, number2))
}

