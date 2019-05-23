package test

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

type T struct {
	A int
	B string
}

func TestAppendValue(t *testing.T) {
	AppendValue()
}

func TestMyReader_Read(t *testing.T) {
	var reader io.Reader
	reader = &MyReader{Name: "a.txt"}
	fmt.Printf("reader: %v\n", reader)
	PrintLine()
	var x float64 = 3.14
	fmt.Printf("type: %v\n", reflect.TypeOf(x))
	PrintLine()
	var y float64 = 3.14
	v := reflect.ValueOf(y)
	fmt.Println("yKind", v.Kind())
	fmt.Println("yType", v.Type())
	fmt.Println("yString", v.String())
	fmt.Println("yValue", v.Float())
	PrintLine()
	p := reflect.ValueOf(&y)
	fmt.Println("pType", p.Type())
	fmt.Println("pCanset", p.CanSet())
	PrintLine()
	m := p.Elem()
	fmt.Println("mCanset", m.CanSet())
	PrintLine()
	m.SetFloat(3.1415926)
	fmt.Println("mValue", m.Interface())
	fmt.Println("mValue", y)
}

func TestWrite(t *testing.T) {
	tt := T{203, "cb203"}
	s := reflect.ValueOf(&tt).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func TestExample(t *testing.T) {
	TypeExample()
	PrintLine()
	ValueExample()
}

func TestScanPerson(t *testing.T) {
	ScanPerson()
}

func PrintLine() {
	fmt.Println("华丽的------------------------分割线")
}
