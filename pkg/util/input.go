package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

func DoEachRowAll(input []byte, f func(row []byte, rows [][]byte, nr, total int) error) (err error) {
	b := bytes.Split(input, []byte{'\n'})
	return DoEachRowAllBytes(b, f)
}

func DoEachRowAllBytes(rows [][]byte, f func(row []byte, rows [][]byte, nr, total int) error) (err error) {
	var (
		total = len(rows)
	)
	for i := 0; i < total; i++ {
		err = f(rows[i], rows, i, total)
		if err != nil {
			return
		}
	}
	return
}

func DoEachRowFile(filename string, rowFunc func(row []byte, nr int) error) (err error) {
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
	return DoEachRowBuff(bufio.NewReader(f), rowFunc)
}

func DoEachRowBuff(r *bufio.Reader, f func(row []byte, nr int) error) (err error) {
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
			err = f(row, i)
			if err != nil {
				return
			}
		}
	}
	return
}

func DoEachRowFileN(filename string, n int, rowFunc func(rows [][]byte, nr int) error) (err error) {
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
	return DoEachRowBuffN(bufio.NewReader(f), n, rowFunc)
}

func DoEachRowBuffN(r *bufio.Reader, n int, f func(rows [][]byte, nr int) error) (err error) {
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
			rows = append(rows, row)
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
