package util

import (
	"iter"
	"math/big"
)

func Combinate[E any](m int, items ...E) iter.Seq[[]E] {
	if m < 0 {
		panic("cannot be negative")
	}
	if m == 0 {
		return func(yield func([]E) bool) {
			return
		}
	}
	switch len(items) {
	case 0:
		return func(yield func([]E) bool) {
			return
		}
	case 1:
		res := Repeat(m, items[0])
		return func(yield func([]E) bool) {
			yield(res)
			return
		}
	}
	if m < 2 {
		res := make([]E, 1)
		return func(yield func([]E) bool) {
			for _, e := range items {
				res[0] = e
				if !yield(res) {
					return
				}
			}
		}
	}
	zeroRes := Repeat(m, items[0])
	n := PowerInt64(int64(len(items)), int64(m))

	breakPoints := make([]int64, m)
	itemLen := int64(len(items))
	b := n / itemLen
	for i := 0; i < m; i++ {
		breakPoints[i] = b
		b /= itemLen
	}
	var (
		k        int
		val, idx int64
	)
	return func(yield func([]E) bool) {
		if !yield(zeroRes) {
			return
		}
		for i := int64(1); i < n; i++ {
			val = i
			for k, b = range breakPoints {
				if val >= b {
					idx = val / b
					zeroRes[k] = items[idx]
					val -= idx * b
				} else {
					zeroRes[k] = items[0]
				}
			}
			if !yield(zeroRes) {
				return
			}
		}
		return
	}
}

func PowerInt64(x, y int64) int64 {
	bx := big.NewInt(x)
	by := big.NewInt(y)
	return bx.Exp(bx, by, nil).Int64()
}
