package util

import (
	"github.com/linusback/aoc/pkg/errorsx"
	"unsafe"
)

const (
	ErrPatternIsNull errorsx.SimpleError = "pattern cannot be null"
)

var AsciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

type PatternFunc func([]byte) bool

type StringPattern string

func (s StringPattern) Pattern(p []byte) bool {
	if len(p) < len(s) {
		return false
	}
	return BytesEqualString(p[:len(s)], string(s))
}

type Patterner interface {
	Pattern([]byte) bool
}

type Tokenizer struct {
	src      []byte
	patterns []Patterner
}

func NewTokenizer(src []byte, p Patterner, patterns ...Patterner) (t *Tokenizer, err error) {
	if p == nil {
		return nil, ErrPatternIsNull
	}
	for _, pe := range patterns {
		if pe == nil {
			return nil, ErrPatternIsNull
		}
	}
	t = new(Tokenizer)
	t.src = src

	return
}

func BytesEqualString(b []byte, s string) bool {
	return ToUnsafeString(b) == s
}

func ToUnsafeString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}

func PtrToUnsafeString(s *[]byte) string {
	return *(*string)(unsafe.Pointer(s))
}
