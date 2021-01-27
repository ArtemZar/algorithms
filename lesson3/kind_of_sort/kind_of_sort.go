package kind_of_sort

// функция сортирует слайс по возрастанию (пузырьковая сортировка)
// на вход принимает слайс значений типа int
//  возвращает отсортированый слайс значений типа int
func BubbleSort(some_slice []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(some_slice)-1; i++ {
			if some_slice[i+1] < some_slice[i] {
				some_slice[i+1], some_slice[i] = some_slice[i], some_slice[i+1]
				swapped = true
			}
		}
	}
	return some_slice
}

// функция сортирует слайс по возрастанию (шейкерная сортировка)
// на вход принимает слайс значений типа int
//  возвращает отсортированый слайс значений типа int
func ShakeSort(some_slice []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(some_slice)-1; i=i+2 {
			if some_slice[i+1] < some_slice[i] {
				some_slice[i+1], some_slice[i] = some_slice[i], some_slice[i+1]
				swapped = true
			}
		}
		for i := len(some_slice)-2; i > 0; i=i-2 {
			if some_slice[i+1] < some_slice[i] {
				some_slice[i+1], some_slice[i] = some_slice[i], some_slice[i+1]
				swapped = true
			}
		}
	}
	return some_slice
}
