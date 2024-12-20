package util

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/linusback/aoc/pkg/errorsx"
	"io"
	"os"
	"slices"
	"strings"
)

var (
	ErrUnsupportedPositionType errorsx.SimpleError = "type is not supported"
)

type (
	RowFunc       func(row []byte, nr int) error
	ExecFunc      func(*bufio.Reader, RowFunc, ...RowFunc) error
	MultiRowFunc  func(row [][]byte, nr int) error
	MultiExecFunc func(*bufio.Reader, int, MultiRowFunc, ...MultiRowFunc) error
)

type UnsignedPos interface {
	~uint16 | ~uint32 | ~uint64
}

type Coordinate interface {
	uint8 | uint16 | uint32
}

type Coordinater[C Coordinate] interface {
	Y() C
	X() C
}

type Pos[U UnsignedPos] interface {
	~uint16 | ~uint32 | ~uint64
	New(y, x int) U
	IsInside(U) bool
}

type PositionMap[P Pos[U], U UnsignedPos, T any] struct {
	Map       []T
	Positions []U
	MaxPos    U
}

func (p PositionMap[P, E, T]) HasInside(pos P) bool {
	return pos.IsInside(p.MaxPos)
}

func (p PositionMap[P, E, T]) Contains(pos E) bool {
	return slices.Contains(p.Positions, pos)
}

func (p PositionMap[P, E, T]) String() string {
	var sb strings.Builder
	sb.WriteString("{\n")
	sb.WriteString(p.MapString())
	sb.WriteString(fmt.Sprintf("\n\tMax: %v", p.MaxPos))
	sb.WriteString(fmt.Sprintf("\n\tPos: %v\n}", p.Positions))
	return sb.String()
}

func (p PositionMap[P, E, T]) MapString() string {
	var (
		sb      strings.Builder
		t       T
		lastPos E
	)
	sb.WriteString("\tMap: ")
	for _, pos := range p.Positions {
		if pos > lastPos+1 {
			sb.WriteString("\n\t     ")
		}
		t = p.Map[pos]
		switch v := any(t).(type) {
		case uint8:
			sb.WriteByte(v)
		default:
			sb.WriteString(fmt.Sprintf("%v", t))
		}
		lastPos = pos
	}
	return sb.String()
}

func Identity[T any](t T) T {
	return t
}

func ToMapOfPositionsByte[P Pos[U], U UnsignedPos](filename string, extra ...RowFunc) (posMap PositionMap[P, U, byte], err error) {
	return ToMapOfPositions[P](filename, Identity, extra...)
}

func ToMapOfPositions[P Pos[U], U UnsignedPos, V any](filename string, transform func(byte) V, extra ...RowFunc) (posMap PositionMap[P, U, V], err error) {
	var (
		y, x int
		row  []byte
		b    byte
		zero P
	)

	//goland:noinspection GoDfaConstantCondition
	//if zero == nil {
	//	panic("zero value should not be nil")
	//}
	data := make([][]byte, 0, 255)
	err = DoEachRowFile(filename, func(row []byte, nr int) error {
		if nr == 0 {
			x = len(row) - 1
		}
		y = nr
		data = append(data, row)
		return nil
	}, extra...)
	if err != nil {
		return posMap, err
	}
	posMap.MaxPos = zero.New(y, x)
	posMap.Map = make([]V, posMap.MaxPos+1)
	var pos U
	for y, row = range data {
		for x, b = range row {
			pos = zero.New(y, x)
			posMap.Positions = append(posMap.Positions, pos)
			posMap.Map[pos] = transform(b)
		}
	}
	return posMap, nil
}

func DoFile(filename string, execFunc ExecFunc, rowFunc RowFunc, extra ...RowFunc) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			err = errors.Join(err, err2)
		}
	}()
	return execFunc(bufio.NewReader(f), rowFunc, extra...)
}

func DoEachRowFile(filename string, rowFunc RowFunc, extra ...RowFunc) (err error) {
	return DoFile(filename, DoEachRowBuff, rowFunc, extra...)
}

func DoEachRowBuff(r *bufio.Reader, f RowFunc, extra ...RowFunc) (err error) {
	var (
		row  []byte
		done bool
	)
	for i := 0; !done; i++ {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
			err = nil
		}
		if len(row) > 0 {
			// removing delim
			row = row[:len(row)-1]
		}
		if len(row) == 0 {
			if len(extra) > 0 {
				f = extra[0]
				extra = extra[:len(extra)-1]
				continue
			}
			return
		}

		err = f(row, i)
		if err != nil {
			return
		}
	}
	return
}

func DoFileN(filename string, n int, execFunc MultiExecFunc, rowFunc MultiRowFunc, extra ...MultiRowFunc) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		err2 := f.Close()
		if err2 != nil {
			err = errors.Join(err, err2)
		}
	}()
	return execFunc(bufio.NewReader(f), n, rowFunc, extra...)
}

func DoEachRowFileN(filename string, n int, rowFunc MultiRowFunc, extra ...MultiRowFunc) (err error) {
	return DoFileN(filename, n, DoEachRowBuffN, rowFunc, extra...)
}

func DoEachRowBuffN(r *bufio.Reader, n int, f MultiRowFunc, extra ...MultiRowFunc) (err error) {
	var (
		done bool
		row  []byte
		rows = make([][]byte, 0, n)
		i    int
	)
	for i = 0; !done; i++ {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
			err = nil
		}
		if len(row) > 0 {
			// removing delim
			row = row[:len(row)-1]
		}
		if len(row) > 0 {
			rows = append(rows, row)
		} else {
			if len(extra) > 0 {
				f = extra[0]
				extra = extra[:len(extra)-1]
				continue
			}
			return
		}
		if len(rows) == n {
			err = f(rows, i)
			if err != nil {
				return
			}
			rows = rows[:0]
		}
	}
	if len(rows) > 0 {
		err = f(rows, i)
		if err != nil {
			return
		}
		rows = rows[:0]
	}

	return
}
