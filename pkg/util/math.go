package util

func Combinate[E any](n int, items ...E) [][]E {
	if n < 2 {
		return [][]E{items}
	}
	return nil
}
