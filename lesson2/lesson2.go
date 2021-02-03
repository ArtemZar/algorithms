package main

import (
	"fmt"
	"math"
)
var exponentiation_value, degree_value float64
var decimal int32
var resalt_binary []int32

func main() {
	//fmt.Print("Введите число для возведения в степень: ")
	//fmt.Scan(&exponentiation_value)
	//fmt.Print("Введите значение степени: ")
	//fmt.Scan(&degree_value)
	//degree(exponentiation_value, degree_value)
	//fmt.Println("Результат вычислений через рекурсию: ", degree_rec(exponentiation_value, degree_value))

	fmt.Print("Введите число десятичной системы счисления: ")
	fmt.Scan(&decimal)
	binary(decimal)



}

// 2. Реализовать функцию возведения числа a в степень b:
//a. Без рекурсии.
func degree (a, b float64)  {
	//через функцию стандартной библиотеки
	fmt.Println ("Результат вычислений через функцию стандартной библиотеки: ", math.Pow(a, b))

	//через цикл
var i, res float64
	res = a
	for i=1; i<b; i++ {
		res = res*a
	}
	fmt.Println("Результат вычислений через цикл: ", res)
}

//b. Рекурсивно.
func degree_rec(a, b float64) float64 {

	if b == 1 {
		return a
	} else {
		a = a*exponentiation_value
			resalt:=degree_rec(a, b-1)
		return resalt
	}

}

// 1. Реализовать функцию перевода чисел из десятичной системы в двоичную, используя рекурсию.

func binary (a int32)   {
	b := a % 2
	c := a / 2
	if c < 1 {
		resalt_binary = append(resalt_binary, b)
		for i  := len(resalt_binary)-1; i>=0; i-- {
			fmt.Print(resalt_binary[i])
		}

	} else {
		resalt_binary = append(resalt_binary, b)
		binary(c)
	}



}


