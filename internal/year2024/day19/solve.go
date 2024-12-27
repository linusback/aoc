package day19

import (
	"bytes"
	"cmp"
	"fmt"
	"github.com/dolthub/swiss"
	"github.com/linusback/aoc/pkg/util"
	"log"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	exampleFile = "./internal/year2024/day19/example.txt"
	inputFile   = "./internal/year2024/day19/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

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

var (
	towelsByChar [256][]pattern
	towels       []pattern
	patterns     []pattern
	towelMap     [826][]pattern
	oneStripeMap [26]uint64
	//knownPattern  = make(map[string]uint64, 19000)
)

func solve(filename string) (solution1, solution2 string, err error) {
	err = util.DoEachRowFile(filename, parseTowels, parsePatterns)
	if err != nil {
		return
	}

	//log.Println("________________________")
	//for i, t := range towelMap {
	//	if len(t) == 0 {
	//		continue
	//	}
	//	mapPattern(keyToString(i), t)
	//	//log.Printf("%c: %s\n", i, ToString(t))
	//}
	//for i, u := range oneStripeMap {
	//	if u == 0 {
	//		continue
	//	}
	//	log.Printf("%c: 1", i+'a')
	//}

	//solution1, solution2 = solveTowels()
	solution1, solution2 = solveTowelsParallel()

	return
}

func mapPattern(s string, p []pattern) {
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

func solveTowels() (solution1, solution2 string) {
	var res1, res2 uint64
	//knownPattern := make(map[string]uint64, 18500)
	knownPattern := swiss.NewMap[string, uint64](18500)
	for _, t := range patterns {
		if ways := canBeMade(t, 0, knownPattern); ways > 0 {
			res1++
			res2 += ways
		}
	}
	//log.Println("len:", len(knownPattern))
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

func solveTowelsParallel() (solution1, solution2 string) {
	const parallel = 16
	var res1, res2 uint64
	ch := consume(parallel)
	wg := new(sync.WaitGroup)
	wg.Add(parallel)
	for i := 0; i < parallel; i++ {
		go func() {
			defer wg.Done()
			knownPattern := swiss.NewMap[string, uint64](2500)
			for t := range ch {
				if ways := canBeMade(t, 0, knownPattern); ways > 0 {
					atomic.AddUint64(&res1, 1)
					atomic.AddUint64(&res2, ways)
				}
			}
			//fmt.Println("len:", knownPattern.Count())
			//fmt.Println("len:", len(knownPattern))

		}()
	}
	wg.Wait()
	//for _, t := range patterns {
	//	if ways := canBeMade(t, 0); ways > 0 {
	//		res1++
	//		res2 += ways
	//	}
	//}
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

func consume(parallel int) <-chan pattern {
	ch := make(chan pattern, parallel)
	go func() {
		for _, p := range patterns {
			ch <- p
		}
		close(ch)
	}()
	return ch
}

func canBeMade(pattern pattern, res uint64, knownPattern *swiss.Map[string, uint64]) uint64 {
	switch len(pattern) {
	case 0:
		return res + 1
	case 1:
		return res + oneStripeMap[pattern[0]-'a']
	}
	var (
		newWays uint64
		ways    uint64
		key     string
		ok      bool
	)
	key = util.ToUnsafeString(pattern)
	if ways, ok = knownPattern.Get(key); ok {
		return res + ways
	}

	if oneStripeMap[pattern[0]-'a'] == 1 {
		ways = canBeMade(pattern[1:], res, knownPattern)
		newWays += ways
		knownPattern.Put(util.ToUnsafeString(pattern[1:]), ways)
		//knownPattern[util.ToUnsafeString(pattern[1:])] = ways
	}

	towelPatters := getTowelMap(pattern[0], pattern[1])
	if len(towelPatters) == 0 {
		knownPattern.Put(util.ToUnsafeString(pattern), 0)
		//knownPattern[util.ToUnsafeString(pattern)] = 0
		return res
	}

	for _, t := range towelPatters {
		if len(pattern) < len(t) {
			break
		}
		if notMatch(pattern, t) {
			continue
		}
		key = util.ToUnsafeString(pattern[len(t):])
		if ways, ok = knownPattern.Get(key); ok {
			newWays += ways
			continue
		}
		ways = canBeMade(pattern[len(t):], res, knownPattern)
		newWays += ways
		knownPattern.Put(key, ways)
	}
	return res + newWays
}

func notMatch(pattern, t pattern) bool {
	for i := 2; i < len(t); i++ {
		if pattern[i] != t[i] {
			return true
		}
	}
	return false
}

func notMatch2(pattern, t pattern) bool {
	if pattern[0] != t[0] {
		return true
	}
	return !bytes.Equal(t, pattern[:len(t)])
}

func notMatch3(pattern, t pattern) bool {
	if pattern[0] != t[0] {
		return true
	}
	return util.ToUnsafeString(t) != util.ToUnsafeString(pattern[:len(t)])
}

func notMatch4(pattern, t pattern) bool {
	if pattern[0] != t[0] {
		return true
	}
	switch len(t) {
	case 1:
		return false
	case 2:
		return pattern[1] != t[1]
	case 3:
		return pattern[1] != t[1] || pattern[2] != t[2]
	default:
		for i := 1; i < len(t); i++ {
			if pattern[i] != t[i] {
				return true
			}
		}
		return false
	}
}

func parsePatterns(row []byte, _ int) error {
	patterns = append(patterns, row)
	return nil
}

func parseTowels(row []byte, _ int) error {
	var (
		start, i int
		b        byte
		towel    pattern
	)
	for i, b = range row {
		if util.AsciiSpace[b] == 1 {
			start++
			continue
		}
		if b == ',' {
			towel = row[start:i]
			towels = append(towels, towel)
			towelsByChar[row[start]] = append(towelsByChar[row[start]], towel)
			addTowelMaps(towel)

			start = i + 1
		}
	}
	if start < i {
		towel = row[start:]
		towels = append(towels, towel)
		towelsByChar[row[start]] = append(towelsByChar[row[start]], towel)
		addTowelMaps(towel)
	}
	slices.SortFunc(towels, patternSort)
	for _, t := range towelsByChar {
		if len(t) == 0 {
			continue
		}
		slices.SortFunc(t, patternSort)
	}
	for _, t := range towelMap {
		if len(t) == 0 {
			continue
		}
		slices.SortFunc(t, patternSort)
	}
	return nil
}

func patternSort(a, b pattern) int {
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return cmp.Compare(string(a), string(b))
}

func getTowelMap(a, b byte) []pattern {
	return towelMap[uint16(a-'a')<<5|uint16(b-'a')]
}

func keyToString(i int) string {
	k := uint16(i)
	const mask = 1<<5 - 1
	return fmt.Sprintf("%c%c", byte(k>>5)+'a', byte(k&mask)+'a')
}

func addTowelMaps(p pattern) {
	switch len(p) {
	case 0:
		return
	case 1:
		oneStripeMap[p[0]-'a'] = 1
		return
	}
	key := uint16(p[0]-'a')<<5 | uint16(p[1]-'a')
	towelMap[key] = append(towelMap[key], p)
}
