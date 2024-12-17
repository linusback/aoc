package position

import (
	"fmt"
	"math"
)

var (
	Directions8 = [...]Pos8{
		New8Negative(-1, 0), // up
		New8Negative(0, 1),  // right
		New8Negative(1, 0),  // down
		New8Negative(0, -1), // left
	}
)

type Pos8 uint16

func New8(y, x uint8) Pos8 {
	return Pos8(x) | Pos8(y)<<8
}

func New8Negative(y, x int8) Pos8 {
	return Pos8(uint8(x)) | Pos8(uint8(y))<<8
}

func (p Pos8) Y() uint8 {
	return uint8(p >> 8)
}

func (p Pos8) X() uint8 {
	return uint8(p & math.MaxUint8)
}

func (p Pos8) IsInside(maxPos Pos8) bool {
	return p <= maxPos && p.X() <= maxPos.X() // -1 should wrap around to an even bigger number.
}

func (p Pos8) String() string {
	return fmt.Sprintf("(%d, %d)", p.Y(), p.X())
}

func (p Pos8) GoString() string {
	return fmt.Sprintf("{Y:%d, X:%d}", p.Y(), p.X())
}

func (p *Pos8) Add(o Pos8) {
	*p = Add8(*p, o)
}

func (p *Pos8) Sub(o Pos8) {
	*p = Sub8(*p, o)
}

func Add8(a, b Pos8) Pos8 {
	return New8(a.Y()+b.Y(), a.X()+b.X())
}

func Sub8(a, b Pos8) Pos8 {
	return New8(a.Y()-b.Y(), a.X()-b.X())
}

type Pos16 uint32

func New16(y, x uint16) Pos16 {
	return Pos16(x) | Pos16(y)<<8
}

func New16Negative(y, x int16) Pos16 {
	return Pos16(uint16(x)) | Pos16(uint16(y))<<8
}

func (p Pos16) Y() uint16 {
	return uint16(p >> 8)
}

func (p Pos16) X() uint16 {
	return uint16(p & math.MaxUint8)
}

func (p Pos16) IsInside(maxPos Pos16) bool {
	return p <= maxPos && p.X() <= maxPos.X() // -1 should wrap around to an even bigger number.
}

func (p Pos16) String() string {
	return fmt.Sprintf("(%d, %d)", p.Y(), p.X())
}

func (p Pos16) GoString() string {
	return fmt.Sprintf("{Y:%d, X:%d}", p.Y(), p.X())
}

func (p *Pos16) Add(o Pos16) {
	*p = Add16(*p, o)
}

func (p *Pos16) Sub(o Pos16) {
	*p = Sub16(*p, o)
}

func Add16(a, b Pos16) Pos16 {
	return New16(a.Y()+b.Y(), a.X()+b.X())
}

func Sub16(a, b Pos16) Pos16 {
	return New16(a.Y()-b.Y(), a.X()-b.X())
}

type Pos32 uint64

func New32(y, x uint32) Pos32 {
	return Pos32(x) | Pos32(y)<<8
}

func New32Negative(y, x int32) Pos32 {
	return Pos32(uint32(x)) | Pos32(uint32(y))<<8
}

func (p Pos32) Y() uint32 {
	return uint32(p >> 8)
}

func (p Pos32) X() uint32 {
	return uint32(p & math.MaxUint8)
}

func (p Pos32) IsInside(maxPos Pos32) bool {
	return p <= maxPos && p.X() <= maxPos.X() // -1 should wrap around to an even bigger number.
}

func (p Pos32) String() string {
	return fmt.Sprintf("(%d, %d)", p.Y(), p.X())
}

func (p Pos32) GoString() string {
	return fmt.Sprintf("{Y:%d, X:%d}", p.Y(), p.X())
}

func (p *Pos32) Add(o Pos32) {
	*p = Add32(*p, o)
}

func (p *Pos32) Sub(o Pos32) {
	*p = Sub32(*p, o)
}

func Add32(a, b Pos32) Pos32 {
	return New32(a.Y()+b.Y(), a.X()+b.X())
}

func Sub32(a, b Pos32) Pos32 {
	return New32(a.Y()-b.Y(), a.X()-b.X())
}
