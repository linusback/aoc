package position

import (
	"fmt"
	"math"
)

type Dir uint

//goland:noinspection GoSnakeCaseUsage
const (
	Dir_Up Dir = iota
	Dir_Right
	Dir_Down
	Dir_Left
	Dir_UpLeft
	Dir_UpRight
	Dir_DownRight
	Dir_DownLeft
)

type Collection[P Position] [4]P

var (
	directions = [...]Dir{
		Dir_Up,
		Dir_Right,
		Dir_Down,
		Dir_Left,
		Dir_UpLeft,
		Dir_UpRight,
		Dir_DownRight,
		Dir_DownLeft,
	}
	DirectionsPos8          Collection[Pos8]  = createDirections[Pos8]()
	DirectionsPos16         Collection[Pos16] = createDirections[Pos16]()
	DirectionsPos32         Collection[Pos32] = createDirections[Pos32]()
	DirectionsPos           Collection[Pos]   = createDirections[Pos]()
	DirectionsDiagonalPos8  Collection[Pos8]  = createDiagonalDirections[Pos8]()
	DirectionsDiagonalPos16 Collection[Pos16] = createDiagonalDirections[Pos16]()
	DirectionsDiagonalPos32 Collection[Pos32] = createDiagonalDirections[Pos32]()
	DirectionsDiagonalPos   Collection[Pos]   = createDiagonalDirections[Pos]()
)

type Position interface {
	Pos8 | Pos16 | Pos32 | Pos
}

type Positioner[P Position] interface {
	IsInside(P) bool
	Add(P) P
	Sub(P) P
	NewDir(y, x int8) P
	New(x, y int) P
}

type PosMap[P Position] struct {
	Map    []P
	MaxPos P
}

func (d Dir) Pos() (y, x int8) {
	switch d {
	case Dir_Up:
		return -1, 0
	case Dir_Right:
		return 0, 1
	case Dir_Down:
		return 1, 0
	case Dir_Left:
		return 0, -1
	case Dir_UpLeft:
		return -1, -1
	case Dir_UpRight:
		return -1, 1
	case Dir_DownRight:
		return 1, 1
	case Dir_DownLeft:
		return 1, -1
	default:
		panic(fmt.Sprintf("direction %d is not implemented", d))
	}
}

func (c Collection[P]) AddTo() {

}

type Pos8 uint16

func New8(y, x uint8) Pos8 {
	return Pos8(x) | Pos8(y)<<8
}

func New8Negative(y, x int8) Pos8 {
	return Pos8(uint8(x)) | Pos8(uint8(y))<<8
}

func (Pos8) NewDir(y, x int8) Pos8 {
	return New8Negative(y, x)
}

func (Pos8) New(y, x int) Pos8 {
	return New8(uint8(y), uint8(x))
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

func (p Pos8) Add(b Pos8) Pos8 {
	return New8(p.Y()+b.Y(), p.X()+b.X())
}

func (p Pos8) Sub(b Pos8) Pos8 {
	return New8(p.Y()-b.Y(), p.X()-b.X())
}

func (p *Pos8) AddSelf(o Pos8) {
	*p = p.Add(o)
}

func (p *Pos8) SubSelf(o Pos8) {
	*p = p.Sub(o)
}

type Pos16 uint32

func New16(y, x uint16) Pos16 {
	return Pos16(x) | Pos16(y)<<8
}

func New16Negative(y, x int16) Pos16 {
	return Pos16(uint16(x)) | Pos16(uint16(y))<<8
}

func (Pos16) NewDir(y, x int8) Pos16 {
	return New16Negative(int16(y), int16(x))
}

func (Pos16) New(y, x int) Pos16 {
	return New16(uint16(y), uint16(x))
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

func (p Pos16) Add(b Pos16) Pos16 {
	return New16(p.Y()+b.Y(), p.X()+b.X())
}

func (p Pos16) Sub(b Pos16) Pos16 {
	return New16(p.Y()-b.Y(), p.X()-b.X())
}

func (p *Pos16) AddSelf(o Pos16) {
	*p = p.Add(o)
}

func (p *Pos16) SubSelf(o Pos16) {
	*p = p.Sub(o)
}

type Pos32 uint64

func New32(y, x uint32) Pos32 {
	return Pos32(x) | Pos32(y)<<8
}

func New32Negative(y, x int32) Pos32 {
	return Pos32(uint32(x)) | Pos32(uint32(y))<<8
}

func (Pos32) NewDir(y, x int8) Pos32 {
	return New32Negative(int32(y), int32(x))
}

func (Pos32) New(y, x int) Pos32 {
	return New32(uint32(y), uint32(x))
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

func (p Pos32) Add(b Pos32) Pos32 {
	return New32(p.Y()+b.Y(), p.X()+b.X())
}

func (p Pos32) Sub(b Pos32) Pos32 {
	return New32(p.Y()-b.Y(), p.X()-b.X())
}

func (p *Pos32) AddSelf(o Pos32) {
	*p = p.Add(o)
}

func (p *Pos32) SubSelf(o Pos32) {
	*p = p.Sub(o)
}

type Pos struct {
	y, x int64
}

func New(y, x int64) Pos {
	return Pos{
		y: y,
		x: x,
	}
}

func NewNegative(y, x int64) Pos {
	return New(y, x)
}

func (Pos) NewDir(y, x int8) Pos {
	return NewNegative(int64(y), int64(x))
}

func (Pos) New(y, x int) Pos {
	return New(int64(y), int64(x))
}

func (p Pos) IsInside(maxPos Pos) bool {
	return 0 <= p.y && p.y <= maxPos.y && 0 <= p.x && p.x <= maxPos.x
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.y, p.x)
}

func (p Pos) GoString() string {
	return fmt.Sprintf("{Y:%d, X:%d}", p.y, p.x)
}

func (p Pos) Add(b Pos) Pos {
	return New(p.y+b.y, p.x+b.x)
}

func (p Pos) Sub(b Pos) Pos {
	return New(p.y-b.y, p.x-b.x)
}

func (p *Pos) AddSelf(o Pos) {
	p.y += o.y
	p.x += o.x
}

func (p *Pos) SubSelf(o Pos) {
	p.y -= o.y
	p.x -= o.x
}

func CreateDirections[E Position](dir ...Dir) (res []E) {
	if len(dir) == 0 {
		return nil
	}
	var (
		e    E
		x, y int8
	)
	p, ok := any(e).(Positioner[E])
	if !ok {
		panic("create directions: Position needs to implement Positioner[Position]")
	}
	res = make([]E, len(dir))
	for i, d := range dir {
		x, y = d.Pos()
		res[i] = p.NewDir(x, y)
	}
	return res
}

func createDirections[E Position]() (res [4]E) {
	var (
		e    E
		x, y int8
	)
	p, ok := any(e).(Positioner[E])
	if !ok {
		panic("create directions: Position needs to implement Positioner[Position]")
	}
	for _, d := range directions[:4] {
		x, y = d.Pos()
		res[d] = p.NewDir(x, y)
	}
	return res
}

func createDiagonalDirections[E Position]() (res [4]E) {
	var (
		e    E
		x, y int8
	)
	p, ok := any(e).(Positioner[E])
	if !ok {
		panic("create directions: Position needs to implement Positioner[Position]")
	}
	for i, d := range directions[4:] {
		x, y = d.Pos()
		res[i] = p.NewDir(x, y)
	}
	return res
}

func createAllDirections[E Position]() (res [8]E) {
	var (
		e    E
		x, y int8
	)
	p, ok := any(e).(Positioner[E])
	if !ok {
		panic("create directions: Position needs to implement Positioner[Position]")
	}
	for _, d := range directions {
		x, y = d.Pos()
		res[d] = p.NewDir(x, y)
	}
	return res
}
