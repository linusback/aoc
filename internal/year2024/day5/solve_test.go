package day5

import "testing"

const inputForTest = "input"

func Test_Day5_Solve(t *testing.T) {
	_, _, err := solve(inputForTest, getComparer1)
	if err != nil {
		t.Error(err)
	}
}

func Benchmark_Day5_Compare1(b *testing.B) {
	b.SetBytes(15556)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := solve(inputForTest, getComparer1)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Day5_Compare2(b *testing.B) {
	b.SetBytes(15556)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := solve(inputForTest, getComparer2)
		if err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Day5_Compare3(b *testing.B) {
	b.SetBytes(15556)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := solve(inputForTest, getComparer3)
		if err != nil {
			b.Error(err)
		}
	}
}
