package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println(max_number())
	fmt.Println(rendom_num())
	average()
}

// 12.Написать функцию нахождения максимального из трех чисел.
func max_number() (int32, error) {
	var a, b, c, res int32
	a, b, c = 5, 7, 9
	switch {
	case a > b && a > c:
		res = a
	case b > a && b > c:
		res = b
	case c > a && c > b:
		res = c
	}
	return res, nil
}

//11.С клавиатуры вводятся числа, пока не будет введен 0. Подсчитать среднее арифметическое всех
//положительных четных чисел, оканчивающихся на 8.
func average() error {
	var myslice []int32
	for true {
		fmt.Println("Введите число")
		var a int32
		_, err := fmt.Scan(&a) // проверяем ошибку соответсвея типу вводимых данных
		if err != nil {
			fmt.Println("Проверьте типы входных параметров")
			os.Exit(1)
		}
		if a == 0 {
			break
		}
		var res int32
		if a > 0 && a%2 == 0 && a%10 == 8 {
			myslice = append(myslice, a)
			for _, num := range myslice {
				res += num
			}
			result := res / int32(len(myslice))
			fmt.Printf("Среднее арифметическое из среза чисел %v: %v \n", myslice, result)
		}

	}
	return nil
}

// 13.Написать функцию, генерирующую случайное число от 1 до 100
func rendom_num() (int, error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	return n, nil
}
