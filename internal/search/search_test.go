package search

import (
	"sort"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		sortedData []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Find mid",
			args: args{
				sortedData: makeSortedSlice(3),
				target:     1,
			},
			want: 1,
		},
		{
			name: "Short source",
			args: args{
				sortedData: makeSortedSlice(2),
				target:     0,
			},
			want: 0,
		},
		{
			name: "Bound",
			args: args{
				sortedData: makeSortedSlice(1),
				target:     0,
			},
			want: 0,
		},
		{
			name: "Not found",
			args: args{
				sortedData: makeSortedSlice(1),
				target:     -1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.sortedData, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Бенчмарк для бинарного поиска
func BenchmarkBinarySearch(b *testing.B) {
	data := makeSortedSlice(1_000_000)
	target := data[len(data)-1] // Ищем последний элемент (худший случай)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(data, target)
	}
}

// Бенчмарк для сравнения с sort.Search
func BenchmarkStdSearch(b *testing.B) {
	data := makeSortedSlice(1_000_000)
	target := data[len(data)-1]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.SearchInts(data, target)
	}
}

func TestBinarySearchPerformance(t *testing.T) {
	sizes := []struct {
		name string
		size int
	}{
		{"100", 100},
		{"10_000", 10_000},
		{"1_000_000", 1_000_000},
		{"10_000_000", 10_000_000},
	}

	for _, sz := range sizes {
		t.Run(sz.name, func(t *testing.T) {
			data := makeSortedSlice(sz.size)
			target := data[len(data)-1] // Худший случай

			start := time.Now()
			for i := 0; i < 1000; i++ {
				BinarySearch(data, target)
			}
			elapsed := time.Since(start)

			t.Logf("Size %d: %d ops in %v (avg %v/op)",
				sz.size, 1000, elapsed, elapsed/1000)
		})
	}
}

// Вспомогательная функция для создания отсортированного слайса
func makeSortedSlice(size int) []int {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	return data
}
