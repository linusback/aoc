package day19

import (
	"bytes"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"strconv"
	"strings"
)

const (
	exampleFile = "./internal/year2024/day19/example.txt"
	inputFile   = "./internal/year2024/day19/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type towel []byte

func (t towel) String() string {
	return util.ToUnsafeString(t)
}

type towelArr []towel

func (t towelArr) String() string {
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

var (
	towels            towelArr
	patterns          towelArr
	impossiblePattern = make(map[string]struct{})
)

func solve(filename string) (solution1, solution2 string, err error) {
	err = util.DoEachRowFile(filename, parseTowels, parsePatterns)
	if err != nil {
		return
	}
	log.Println("towels", towels)
	log.Println("patterns", patterns)
	var res1 uint64
	for i, t := range patterns {
		log.Printf("testing pattern (%d/%d): %v", i+1, len(patterns), t)
		if canBeMade(t) {
			res1++
			log.Println("can be made", t)
		} else {
			log.Println("NO POSSIBLE", t)
		}
		log.Println("________________")
	}
	log.Println("can make", res1, "patterns")
	solution1 = strconv.FormatUint(res1, 10)
	return
}

func canBeMade(pattern towel) bool {
	if len(pattern) == 0 {
		return true
	}
	for _, t := range towels {
		if !matchesStart(pattern, t) {
			continue
		}
		if _, ok := impossiblePattern[util.ToUnsafeString(pattern[len(t):])]; ok {
			continue
		}
		if canBeMade(pattern[len(t):]) {
			//log.Println("using towel", t)
			return true
		} else {
			impossiblePattern[util.ToUnsafeString(pattern[len(t):])] = struct{}{}
		}
	}
	return false
}

func matchesStart(pattern, t towel) bool {
	if len(pattern) < len(t) {
		return false
	}
	return bytes.Equal(pattern[:len(t)], t)
}

func parsePatterns(row []byte, _ int) error {
	patterns = append(patterns, row)
	return nil
}

func parseTowels(row []byte, _ int) error {
	var (
		start, i int
		b        byte
	)
	for i, b = range row {
		if util.AsciiSpace[b] == 1 {
			start++
			continue
		}
		if b == ',' {
			towels = append(towels, row[start:i])
			start = i + 1
		}
	}
	if start < i {
		towels = append(towels, row[start:])
	}
	return nil
}
