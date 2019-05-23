package test

import (
	"fmt"
	"reflect"
	"unsafe"
)

type user struct {
	b byte
	i int32
	j int64
}

func (u user) SayHi(name string) {
	fmt.Println("Hello,", name)
}

func Say() {
	var u user
	u.SayHi("XuHao")
}

func SizeOf() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic:", err)
		}
	}()
	fmt.Println("true", unsafe.Sizeof(true))
	fmt.Println("int8(0)", unsafe.Sizeof(int8(0)))
	fmt.Println("int16(10)", unsafe.Sizeof(int16(10)))
	fmt.Println("int32", unsafe.Sizeof(int32(1000000000)))
	fmt.Println("int64", unsafe.Sizeof(int64(10000000000000000)))
	fmt.Println("int", unsafe.Sizeof(int(1000000000000000)))
}

func AlignOf() {
	var b bool
	var i8 int8
	var i16 int16
	var f32 float32
	var f64 float64
	var s string
	var m map[string]string
	var p *int32

	fmt.Println("b", unsafe.Alignof(b))
	fmt.Println("i8", unsafe.Alignof(i8))
	fmt.Println("i16", unsafe.Alignof(i16))
	fmt.Println("f32", unsafe.Alignof(f32))
	fmt.Println("f64", unsafe.Alignof(f64))
	fmt.Println("s", unsafe.Alignof(s))
	fmt.Println("m", unsafe.Alignof(m))
	fmt.Println("p", unsafe.Alignof(p))

	var x string
	fmt.Println("x", reflect.TypeOf(x).Align())
}

func OffsetOf() {
	var u user
	offsetJ := unsafe.Offsetof(u.j)
	offsetI := unsafe.Offsetof(u.i)
	offsetB := unsafe.Offsetof(u.b)

	offsetBB := reflect.TypeOf(u).Field(0).Offset
	offsetII := reflect.TypeOf(u).Field(1).Offset
	offsetJJ := reflect.TypeOf(u).Field(2).Offset

	fmt.Println("offset b", offsetB, offsetBB)
	fmt.Println("offset i", offsetI, offsetII)
	fmt.Println("offset j", offsetJ, offsetJJ)
}
