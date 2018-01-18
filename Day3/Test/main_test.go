package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(2, 4)
	if total != 6 {
		t.Errorf("Sum was incorrect. Got: %d, wanted: %d", total, 6)
	}
}

func TestSumTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 1, 3},
		{2, 2, 4},
		{3, 2, 5},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum (%d + %d) was incorrect. Got: %d, wanted: %d", table.x, table.y, total, table.n)
		}
	}
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFibN1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFibN2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFibN3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFibN10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFibN20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFibN30(b *testing.B) { benchmarkFib(30, b) }

func ExampleReverse() {
	fmt.Println(Reverse("Hello"))
	// Output: olleH
}
