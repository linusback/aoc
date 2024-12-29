package day19

import (
	"fmt"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"slices"
	"strings"
)

type pattern []byte

func (t pattern) String() string {
	return util.ToUnsafeString(t)
}

func ToString(t []pattern) string {
	if len(t) == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteByte('[')
	sb.Write(t[0])
	for _, to := range t[1:] {
		sb.WriteByte(' ')
		sb.Write(to)
	}
	sb.WriteByte(']')
	return sb.String()
}

func PrintPattern(s string, p []pattern) {
	var arr [255][]pattern
	for _, p2 := range p {
		if len(p2) < 2 {
			continue
		}
		arr[p2[1]] = append(arr[p2[1]], p2)
	}
	fmt.Println("byte(s): ", s)
	for i, p2 := range arr {
		if len(p2) == 0 {
			continue
		}
		slices.SortFunc(p2, patternSort)
		log.Printf("%c: %s\n", i, ToString(p2))
	}
}
