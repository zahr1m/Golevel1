package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, n!): ")
	fmt.Scanln(&op)
	if op == "n!" {
		fmt.Printf("Результат выполнения операции:  %.2f\n ", factorial(int(a)))
	} else {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				fmt.Println("На 0 делить нельзя!")
			} else {
				res = a / b
			}
		case "^":
			res = math.Pow(a, b)
		default:
			fmt.Println("Операция выбрана неверно")
			os.Exit(1)
			fmt.Printf("Результат выполнения операции:  %.2f\n ", res)
		}
	}
}

func factorial(n int) float64 {
	result := 1.0
	for i := 1; i <= n; i++ {
		result *= float64(i)
	}
	return result
}