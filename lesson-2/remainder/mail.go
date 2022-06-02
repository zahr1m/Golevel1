package main

import "fmt"

func main()  {
	var a uint16
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&a)
	b := a/100 % 100
	c := a/10 % 10
	d := a % 10
	fmt.Println("Сотен:",b)
	fmt.Println("Десятков:",c)
	fmt.Println("Едениц:",d)
}
