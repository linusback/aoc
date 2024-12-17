package day9

import (
	"bytes"
	"github.com/linusback/aoc/pkg/util"
	"os"
	"strconv"
)

const (
	exampleFile = "./internal/year2024/day9/example.txt"
	inputFile   = "./internal/year2024/day9/input.txt"
)

func Solve() (solution1, solution2 string, err error) {
	return solve(inputFile)
}

type file struct {
	start int
	size  []byte
}

func solve(filename string) (solution1, solution2 string, err error) {
	byteArr, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	byteArr = bytes.TrimSpace(byteArr)

	//b := generatePrintRepresentation(byteArr)
	arr := generateInt64(byteArr)
	arr2 := make([]int64, len(arr))
	copy(arr2, arr)
	//log.Printf("%v\n", arr)
	left := int64(0)
	right := int64(len(arr) - 1)
	for ; left < right; left++ {
		if arr[left] > -1 {
			continue
		}
		for ; right > left; right-- {
			if arr[right] == -1 {
				continue
			}
			arr[left] = arr[right]
			arr[right] = -1
			break
		}
	}
	//log.Printf("%v\n", arr)
	var acc int64
	for i := int64(0); i <= left; i++ {
		if arr[i] == -1 {
			break
		}
		acc += i * arr[i]
	}
	solution1 = strconv.FormatInt(acc, 10)

	// TODO should probably just create some sort of type
	arrLen := int64(len(arr) - 1)

	var sizeFree, sizeFile int64
	right = arrLen
	for right > 0 {
		if arr2[right] == -1 {
			right--
			continue
		}
		sizeFile = getFileSpace(arr2[:right+1])
		for left = 0; left < right; {
			if arr2[left] > -1 {
				left++
				continue
			}
			sizeFree = getFreeSpace(arr2[left:])

			//log.Println("free: ", sizeFree, " file:", sizeFile)
			//log.Println("free: ", arr2[left:left+sizeFree])
			//log.Println("file: ", arr2[right+1-sizeFile:right+1])
			if sizeFree >= sizeFile {
				//log.Println("swap: ", arr2[left:left+sizeFree], sizeFree, " <-> ", arr2[right+1-sizeFile:right+1])
				switchArr(arr2[left:left+sizeFree], arr2[right+1-sizeFile:right+1])
				//right -= sizeFile - 1
				break
			}
			left += sizeFree
			//left += sizeFree - 1

		}
		right -= sizeFile

	}
	//log.Println("setting array to expected answer")
	//arr2 = []int64{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}
	//log.Printf("%v\n", arr2)
	acc = 0
	for i := int64(0); i <= arrLen; i++ {
		if arr2[i] == -1 {
			continue
		}
		acc += i * arr2[i]
	}
	//log.Println("res 2:", acc)
	solution2 = strconv.FormatInt(acc, 10)
	return
}

func switchArr(arr1, arr2 []int64) {
	if len(arr1) < len(arr2) {
		panic("length of arr 1 and arr 2 are different")
	}
	for i := range len(arr2) {
		arr1[i] = arr2[i]
		arr2[i] = -1
	}
}
func getFreeSpace(arr []int64) int64 {
	var idx int64
	for ; idx < int64(len(arr)); idx++ {
		if arr[idx] != -1 {
			break
		}
	}
	return idx
}

func getFileSpace(arr []int64) int64 {
	var size int64
	val := arr[len(arr)-1]

	for k := len(arr) - 1; k >= 0; k-- {
		if arr[k] != val {
			break
		}
		size++
	}
	return size
}

func generateInt64(arr []byte) (res []int64) {
	var idx, val int64
	for i, b := range arr {
		b -= '0'
		if i%2 == 0 {
			val = idx
			idx++
		} else {
			val = -1
		}
		res = util.AppendRepeat(res, int(b), val)
	}
	return res
}

func generatePrintRepresentation(arr []byte) (res []byte) {
	var idx int64
	slice := make([]byte, 0, 16)
	for i, b := range arr {
		b -= '0'
		slice = slice[:0]
		if i%2 == 0 {
			slice = strconv.AppendInt(slice, idx, 10)
			idx++
		} else {
			slice = append(slice, '.')
		}
		res = appendRepeat(res, slice, b)
	}
	return res
}

func appendRepeat(res, msg []byte, n byte) []byte {
	for range n {
		res = append(res, msg...)
	}
	return res
}
