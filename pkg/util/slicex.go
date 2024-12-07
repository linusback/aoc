package util

import "slices"

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
