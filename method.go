package test

import "fmt"

type Meower interface {
	Meow()
}

type Monkey struct {
	Name string
	Color string
}

type SmallMonkey struct {
	Monkey
	Size int
}

func Greet(meower Meower) {
	meower.Meow()
}

func (m Monkey) Meow() {
	fmt.Println("Name:", m.Name, "Color:", m.Color)
}

//func (m *Monkey) Meow() {
//	fmt.Println("Name:", m.Name, "Color:", m.Color)
//}

func DoMeow() {
	m := Monkey{"chenyubo", "black"}
	Greet(m)

	mm := &Monkey{"chenhongyuan", "white"}
	Greet(mm)

	sm := SmallMonkey{m, 10}
	fmt.Println("small monkey name:", sm.Name, "small monkey size:", sm.Size)
}