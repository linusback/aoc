package position

import (
	"fmt"
	"log"
	"testing"
)

func Test_Pos8_Add(t *testing.T) {
	p1 := New8(3, 6)
	p2 := New8Negative(3, -4)
	p1.Add(p2)
	fmt.Println(p1)
}

func Test_Pos8_String(t *testing.T) {
	type dummy struct {
		Y, X uint8
	}
	d := dummy{
		Y: 7,
		X: 3,
	}
	fmt.Printf("%v\n", d)
	fmt.Printf("%+v\n", d)
	fmt.Printf("%#v\n", d)

	p := New8(7, 3)

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
}

func TestPos_AddSelf(t *testing.T) {
	p := New(1, 1)
	p.AddSelf(New(2, 2))
	log.Println(p)
}
