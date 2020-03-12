package test

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func DoNewBuffer() {
	var b1 bytes.Buffer
	var b2 = new(bytes.Buffer)
	var b3 = bytes.NewBuffer([]byte{})
	var b4 = bytes.NewBufferString("geek")

	b1.Write([]byte("geek"))
	b2.Write([]byte{'g', 'e', 'e', 'k'})
	binary.Write(b3, binary.BigEndian, int32(100))

	fmt.Printf("%T, %v\n", b1, b1.String())
	fmt.Printf("%T, %v\n", b2, b2)
	fmt.Printf("%T, %v\n", b3, b3.Bytes())
	fmt.Printf("%T, %v\n", b4, b4)

	var l = make([]byte, 3)
	b4.Read(l)
	fmt.Println(string(l), b4.String())
}
