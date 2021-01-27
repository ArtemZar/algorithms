package main

import (
	"fmt"
	"github.com/ArtemZar/algorithms/lesson3/kind_of_sort"
	"math/rand"
	"time"
)



// Функция createRandomSlice создает произвольный срез для дальнейшей сортировки
// значения на вход не принимаются
// возвращает слайс значений типа int
func createRandomSlice() []int {
	count := 5
	sliceForSort := make([]int, 0, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(1000)
		sliceForSort = append(sliceForSort, n)
	}
	return sliceForSort
}

func main() {
	sliceForBubbleSort := createRandomSlice()
	fmt.Println("Исходный слайс: ", sliceForBubbleSort)
	fmt.Println("Алгоритм сортировки пузырьком: ", kind_of_sort.BubbleSort(sliceForBubbleSort))

	sliceForShakeSort := createRandomSlice()
	fmt.Println("Исходный слайс: ", sliceForShakeSort)
	fmt.Println("Алгоритм сортировки пузырьком: ", kind_of_sort.ShakeSort(sliceForShakeSort))

}



