package day19

import (
	"github.com/linusback/aoc/pkg/util"
	"iter"
	"strconv"
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

var (
	patterns = make([]pattern, 0, 400)
	//towels          = make([]pattern, 450)
	trie = make([]uint16, 4800)
	//trie2 [6][]uint16
	tLen      uint16 = 6
	transform util.TransformRowFunc
)

func solve(filename string) (solution1, solution2 string, err error) {
	err = util.DoEachRowFile(filename, parseTowels, parsePatterns)
	if err != nil {
		return
	}

	//transform = transformTowels
	//err = util.DoEachByteFile(filename, TransformTowel)
	//if err != nil {
	//	return
	//}

	//log.Println("trie", trie)
	//log.Println("patterns len", len(patterns))

	//log.Println("max len", maxlen)
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
	solution1, solution2 = solveTowels()

	return
}

func solveTowels() (solution1, solution2 string) {
	var res1, res2 uint64
	//knownPattern := make(map[string]uint64, 18500)
	knownPattern := make([]int64, 61)
	for _, t := range patterns {
		if ways := canBeMadeNew(t, knownPattern); ways > 0 {
			res1++
			res2 += ways
		}
	}
	//log.Println("len:", len(knownPattern))
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

func solveTowelsParallel() (solution1, solution2 string) {
	const parallel = 16
	var (
		res1, res2 uint64
		n          uint32
	)
	wg := new(sync.WaitGroup)
	wg.Add(parallel)
	pLen := uint32(len(patterns))
	for i := 0; i < parallel; i++ {
		go func() {
			defer wg.Done()
			knownPattern := make([]int64, 61)
			var t pattern
			var ways uint64
			for k := atomic.AddUint32(&n, 1) - 1; k < pLen; k = atomic.AddUint32(&n, 1) - 1 {
				t = patterns[k]
				//knownPattern = resetKnownPattern(knownPattern)
				//if ways = canBeMade(t, 0, knownPattern); ways > 0 {
				//	atomic.AddUint64(&res1, 1)
				//	atomic.AddUint64(&res2, ways)
				//}
				if ways = canBeMadeNew(t, knownPattern); ways > 0 {
					atomic.AddUint64(&res1, 1)
					atomic.AddUint64(&res2, ways)
				}
			}

		}()
	}
	wg.Wait()
	return strconv.FormatUint(res1, 10), strconv.FormatUint(res2, 10)
}

// canBeMadeNew is adapted from https://github.com/maneatingape/advent-of-code-rust/blob/main/src/year2024/day19.rs
func canBeMadeNew(p pattern, knownPattern []int64) uint64 {
	var i uint16

	size := len(p)
	knownPattern = knownPattern[:size+1]
	clear(knownPattern)

	knownPattern[0] = 1
	for start := 0; start < size; start++ {
		if knownPattern[start] > 0 {
			i = 0
			for end := start; end < size; end++ {
				i = trie[i+uint16(p[end])]
				if i == 0 {
					break
				}
				knownPattern[end+1] += int64(trie[i+3]) * knownPattern[start]
			}
		}
	}
	//log.Printf("pattern %v\n\t%v\n", p, knownPattern)
	return uint64(knownPattern[size])
}

func parsePatterns(row []byte, _ int) error {
	hashBytes(row)
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
			hashBytes(towel)
			setTrieTowel(towel)
			start = i + 1
		}
	}
	if start < i {
		towel = row[start:]
		hashBytes(towel)
		setTrieTowel(towel)
	}

	//slices.SortFunc(towels, patternSortPreHash)
	//for _, t := range towels {
	//	setTrieTowel(t)
	//}

	//log.Println(trie)
	//log.Printf("trie len %d, cap %d\n", len(trie), cap(trie))
	//log.Println(math.MaxUint16)
	//log.Println("len", len(towels))
	return nil
}

func TransformTowel(b byte) {
	transform(b)
}

var towelBuff = make(pattern, 0, 8)

func transformTowels(b byte) {
	switch b {
	case 'b', 'g', 'r', 'u', 'w':
		towelBuff = append(towelBuff, perfectHashByte(b))
	case ',':
		setTrieTowel(towelBuff)
		towelBuff = towelBuff[:0]
	case '\n':
		setTrieTowel(towelBuff)
		transform = transformEmptyRow
	}
}
func transformEmptyRow(_ byte) {
	transform = transformPatterns
}

var patternBuff = make(pattern, 0, 61)

func transformPatterns(b byte) {
	switch b {
	case '\n':
		if len(patternBuff) == 0 {
			return
		}
		p := make(pattern, len(patternBuff))
		copy(p, patternBuff)
		patternBuff = patternBuff[:0]
		patterns = append(patterns, p)
	default:
		patternBuff = append(patternBuff, perfectHashByte(b))
	}
}

// perfectHash is adapted from https://github.com/maneatingape/advent-of-code-rust/blob/main/src/year2024/day19.rs
// Hashes the five possible color values white (w), blue (u), black (b), red (r), or green (g)
// to 0, 2, 4, 5 and 1 respectively. This compresses the range to fit into an array of 6 elements.
func perfectHash(b byte) uint16 {
	return uint16((b ^ (b >> 4)) & 7)
}

func perfectHashByte(b byte) byte {
	return (b ^ (b >> 4)) & 7
}

func hashBytes(bArr []byte) {
	for i, b := range bArr {
		bArr[i] = perfectHashByte(b)
	}
}

func setTrieTowel(towel pattern) {
	var i, j uint16
	//log.Printf("setting towel %v\n", towel)
	for _, b := range towel {

		j = uint16(b)
		//log.Printf("hash: %c -> %d\n", b, j)
		//log.Printf("%d + %d = %d, %d", i, j, i+j, len(trie))
		if trie[i+j] == 0 {
			trie[i+j] = tLen
			i = tLen
			tLen += 6
			//append(trie, 0, 0, 0, 0, 0, 0)
		} else {
			i = trie[i+j]
		}
	}
	trie[i+3] = 1
}

func patternSortPerfectHash(a, b pattern) int {
	var ha, hb uint16
	for i := 0; i < len(a) && i < len(b); i++ {
		ha, hb = perfectHash(a[i]), perfectHash(b[i])
		if ha < hb {
			return -1
		}
		if ha > hb {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func patternSortPreHash(a, b pattern) int {
	//var ha, hb byte
	for i := 0; i < len(a) && i < len(b); i++ {
		ha, hb := a[i], b[i]
		if ha < hb {
			return -1
		}
		if ha > hb {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}

func matches(p pattern) iter.Seq[int] {
	tLen := min(len(p), 8)
	var (
		trieI uint16
		res   int
	)

	return func(yield func(int) bool) {
		for i := 0; i < tLen; i++ {
			res++
			trieI = trie[trieI+perfectHash(p[i])]
			if trieI == 0 {
				break
			}
			if trie[trieI+3] == 1 {
				if !yield(res) {
					break
				}
			}
		}
	}
}

func matchesGet(p pattern) iter.Seq[pattern] {
	tLen := min(len(p), 8)
	towel := make(pattern, 0, 8)
	var trieI uint16
	//log.Println("initial:", trie[trieI:trieI+6])
	return func(yield func(pattern) bool) {
		for i := 0; i < tLen; i++ {
			towel = append(towel, p[i])
			//log.Printf("i: %d, hash: %d\n", trieI, perfectHash(p[i]))
			//log.Println("current:", trie[trieI:trieI+6])
			trieI = trie[trieI+perfectHash(p[i])]
			if trieI == 0 {
				break
			}
			if trie[trieI+3] == 1 {
				if !yield(towel) {
					break
				}
			}
		}
	}
}
