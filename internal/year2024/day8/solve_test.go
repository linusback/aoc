package day8

import "testing"

func Test_utf(t *testing.T) {
	s := "...#....0..."
	for i, r := range s {
		t.Logf("%d: %c (%d)\n", i, r, r)
	}
}
