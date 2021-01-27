package kind_of_sort

import (
	"testing"

)

var result []int

func BenchmarkBubbleSort(b *testing.B) {
	var res []int
	abc := []int{484, 935, 639, 140, 228, 345, 222, 987, 444, 12, 77, 456}
	for i := 0; i < b.N; i++ {
		res = BubbleSort(abc)
	}
	result = res
}

func BenchmarkShakeSort(b *testing.B) {
	var res []int
	abc := []int{484, 935, 639, 140, 228, 345, 222, 987, 444, 12, 77, 456}
	for i := 0; i < b.N; i++ {
		res = BubbleSort(abc)
	}
	result = res
}
