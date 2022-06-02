package main

import "fmt"

func main()  {
	var a int
	var b int
	fmt.Print("Введите сторону a: ")
	fmt.Scanln(&a)
	fmt.Print("Введите сторону b: ")
	fmt.Scanln(&b)
	S := a*b
	fmt.Println("Площадь прямоугольника равна:", S)
}