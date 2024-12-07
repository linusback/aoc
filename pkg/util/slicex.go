package util

import "slices"

func AppendUnique[S interface{ ~[]E }, E comparable](s S, e ...E) S {
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
