package util

import "fmt"

var (
	IsNumber         = [256]uint8{'0': 1, '1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1}
	IsLetter         = [256]uint8{'A': 1, 'B': 1, 'C': 1, 'D': 1, 'E': 1, 'F': 1, 'G': 1, 'H': 1, 'I': 1, 'J': 1, 'K': 1, 'L': 1, 'M': 1, 'N': 1, 'O': 1, 'P': 1, 'Q': 1, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 1, 'W': 1, 'X': 1, 'Y': 1, 'Z': 1, 'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1, 'l': 1, 'm': 1, 'n': 1, 'o': 1, 'p': 1, 'q': 1, 'r': 1, 's': 1, 't': 1, 'u': 1, 'v': 1, 'w': 1, 'x': 1, 'y': 1, 'z': 1}
	IsUpperLetter    = [256]uint8{'A': 1, 'B': 1, 'C': 1, 'D': 1, 'E': 1, 'F': 1, 'G': 1, 'H': 1, 'I': 1, 'J': 1, 'K': 1, 'L': 1, 'M': 1, 'N': 1, 'O': 1, 'P': 1, 'Q': 1, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 1, 'W': 1, 'X': 1, 'Y': 1, 'Z': 1}
	IsLowerLetter    = [256]uint8{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1, 'l': 1, 'm': 1, 'n': 1, 'o': 1, 'p': 1, 'q': 1, 'r': 1, 's': 1, 't': 1, 'u': 1, 'v': 1, 'w': 1, 'x': 1, 'y': 1, 'z': 1}
	IsAlphaNumerical = [256]uint8{'0': 1, '1': 1, '2': 1, '3': 1, '4': 1, '5': 1, '6': 1, '7': 1, '8': 1, '9': 1, 'A': 1, 'B': 1, 'C': 1, 'D': 1, 'E': 1, 'F': 1, 'G': 1, 'H': 1, 'I': 1, 'J': 1, 'K': 1, 'L': 1, 'M': 1, 'N': 1, 'O': 1, 'P': 1, 'Q': 1, 'R': 1, 'S': 1, 'T': 1, 'U': 1, 'V': 1, 'W': 1, 'X': 1, 'Y': 1, 'Z': 1, 'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1, 'f': 1, 'g': 1, 'h': 1, 'i': 1, 'j': 1, 'k': 1, 'l': 1, 'm': 1, 'n': 1, 'o': 1, 'p': 1, 'q': 1, 'r': 1, 's': 1, 't': 1, 'u': 1, 'v': 1, 'w': 1, 'x': 1, 'y': 1, 'z': 1}
)

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func ParseInt64(arr []byte) (res int64, err error) {
	var mult int64 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			err = fmt.Errorf("failed to parse %s to int", string(arr))
			return
		}
		res += int64(arr[i]-'0') * mult
		mult *= 10
	}
	return
}

func ParseInt64NoError(arr []byte) (res int64) {
	var mult int64 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == '-' {
			return -res
		}
		if arr[i] < '0' || arr[i] > '9' {
			return
		}
		res += int64(arr[i]-'0') * mult
		mult *= 10
	}
	return
}

func ParseInt64ArrNoError(arr []byte) (res []int64) {
	start := -1
	for i := 0; i < len(arr); i++ {
		if start == -1 && ('0' <= arr[i] && arr[i] <= '9') || arr[i] == '-' {
			start = i
			continue
		}
		if start > -1 && ('0' > arr[i] || arr[i] > '9') {
			res = append(res, ParseInt64NoError(arr[start:i]))
			start = -1
		}
	}
	if start > -1 {
		res = append(res, ParseInt64NoError(arr[start:]))
	}

	return res
}

func ParseInt[T Signed](arr []byte) (res T) {
	var mult T = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == '-' {
			return -res
		}
		if arr[i] < '0' || arr[i] > '9' {
			return
		}
		res += T(arr[i]-'0') * mult
		mult *= 10
	}
	return
}

func ParseIntArr[T Signed](arr []byte) (res []T) {
	start := -1
	for i := 0; i < len(arr); i++ {
		if start == -1 && ('0' <= arr[i] && arr[i] <= '9') || arr[i] == '-' {
			start = i
			continue
		}
		if start > -1 && ('0' > arr[i] || arr[i] > '9') {
			res = append(res, ParseInt[T](arr[start:i]))
			start = -1
		}
	}
	if start > -1 {
		res = append(res, ParseInt[T](arr[start:]))
	}

	return res
}

func ParseInt64ArrNoErrorCache(arr []byte, input []int64) (res []int64) {
	start := -1
	res = input[:0]
	for i := 0; i < len(arr); i++ {
		if start == -1 && ('0' <= arr[i] && arr[i] <= '9') || arr[i] == '-' {
			start = i
		}
		if start > -1 && ('0' > arr[i] || arr[i] > '9') {
			res = append(res, ParseInt64NoError(arr[start:i]))
			start = -1
		}
	}
	if start > -1 {
		res = append(res, ParseInt64NoError(arr[start:]))
	}

	return res
}

func ParseUint8NoError(arr []byte) (res uint8) {
	var mult uint8 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			return
		}
		res += (arr[i] - '0') * mult
		mult *= 10
	}
	return
}

func ParseUint64ArrNoError(arr []byte) (res []uint64) {
	start := -1
	for i := 0; i < len(arr); i++ {
		if start == -1 && '0' <= arr[i] && arr[i] <= '9' {
			start = i
		}
		if start > -1 && ('0' > arr[i] || arr[i] > '9') {
			res = append(res, ParseUint64NoError(arr[start:i]))
			start = -1
		}
	}
	if start > -1 {
		res = append(res, ParseUint64NoError(arr[start:]))
	}

	return res
}

func ParseUint64NoError(arr []byte) (res uint64) {
	var mult uint64 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			return
		}
		res += uint64(arr[i]-'0') * mult
		mult *= 10
	}
	return
}

func ParseUint64IgnoreAll(arr []byte) (res uint64) {
	var mult uint64 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			continue
		}
		res += uint64(arr[i]-'0') * mult
		mult *= 10
	}
	return
}
