package util

import (
	"iter"
	"slices"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Repeat[I Integer, E any](n I, e E) []E {
	if n < 0 {
		panic("cannot be negative")
	}
	res := make([]E, n)
	for i := I(0); i < n; i++ {
		res[i] = e
	}
	return res
}

func AppendRepeat[S ~[]E, E any, I Integer](s S, n I, e E) []E {
	if n < 0 {
		panic("cannot be negative")
	}
	for i := I(0); i < n; i++ {
		s = append(s, e)
	}
	return s
}

func ToKeysSeq2[S ~[]K, K comparable, V any](s S, v V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, k := range s {
			if !yield(k, v) {
				return
			}
		}
	}
}

func CountFunc[S ~[]E, E any](s S, f func(E) bool) (res uint64) {
	for _, e := range s {
		if f(e) {
			res++
		}
	}
	return res
}

func AppendUnique[S ~[]E, E comparable](s S, e ...E) S {
	toAppend := make([]E, 0, len(e))
	for _, c := range e {
		if !slices.Contains(s, c) {
			toAppend = append(toAppend, c)
		}
	}
	if len(toAppend) == 0 {
		return s
	}
	return append(s, toAppend...)
}

func Unique[S ~[]E, E comparable](s S) (res S) {
	res = make(S, 0, len(s))
	for _, e := range s {
		if !slices.Contains(res, e) {
			res = append(res, e)
		}
	}
	return res
}

func LenUnique[S ~[]E, E comparable](s S) (res int) {
	for i, e := range s {
		if !slices.Contains(s[:i], e) {
			res++
		}
	}
	return res
}

func AppendUniqueFunc[S ~[]E, E comparable](s S, cmp func(E) bool, e ...E) S {
	switch len(e) {
	case 0:
		return s
	case 1:
		if !slices.ContainsFunc(s, cmp) {
			return append(s, e[0])
		}
		return s
	}
	toAppend := make([]E, 0, len(e))
	for _, c := range e {
		if !slices.ContainsFunc(s, cmp) {
			toAppend = append(toAppend, c)
		}
	}
	if len(toAppend) == 0 {
		return s
	}
	return append(s, toAppend...)
}
