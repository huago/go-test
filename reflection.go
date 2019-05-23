package test

import (
	"fmt"
	"os"
	"reflect"
)

type MyReader struct {
	Name string
}

type Person struct {
	Name    string
	Sex     string
	Age     int
	Address string
	School  string
	Company string
	Wechat  string
	Phone   string
}

func CopyStruct(src, dest interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dest).Elem()

	stype := reflect.TypeOf(src).Elem()
	for j := 0; j < stype.NumField(); j++ {
		typ := stype.Field(j).Name
		dval2 := dval.FieldByName(typ)
		dval2.Set(reflect.Value{})
	}

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}

		dvalue.Set(value)
	}
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

func (r MyReader) Read(p []byte) (n int, err error) {
	n = 100
	err = nil
	return
}

func Write() {

}

func AppendValue() {
	var person Person
	person.Name = "chenyubo"
	person.Sex = "nan"
	person.Age = 12
	person.Address = "beijing"
	person.School = "cup"
	person.Company = "didi"
	person.Wechat = "cup_chenyubo"
	person.Phone = "18810114921"



	personMap := make(map[string]interface{})
	pValue := reflect.ValueOf(person)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Error happened! %v\n", err)
		}
	}()

	pType := reflect.TypeOf(person)
	for i := 0; i < pType.NumField(); i++ {
		personMap[pType.Field(i).Name] = pValue.Field(i).Interface()
	}

	fmt.Printf("%v\n", personMap)
}

func TypeExample() {
	mr := MyReader{Name: "chenyubo"}

	t := reflect.TypeOf(mr)

	fmt.Println(t)

	var i interface{} = 1
	var j interface{} = new(interface{})
	j = 1234
	var z interface{} = new(float64)
	z = 12.34
	var x interface{} = new(interface{})
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(reflect.TypeOf(x))
}

func ValueExample() {
	var a interface{} = new(int)
	a = 100
	//v是Value类型
	v := reflect.ValueOf(a)

	aa := v.Interface().(interface{})
	fmt.Printf("%v\n", aa)

	typ := v.Type()
	fmt.Println(typ)
}

func ScanPerson() {
	p := Person{
		Name:    "chenyubo",
		Sex:     "man",
		School:  "cup",
		Address: "Beijing Haidian",
		Phone:   "18810114921",
		Wechat:  "cup_chenyubo",
		Company: "didi",
		Age:     24,
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error Happened!", err)
			os.Exit(0)
		}
	}()

	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("Name", t.Field(i).Name, "Type", t.Field(i).Type)
	}

	fmt.Println("-----------------------", t.NumMethod())

	for j := 0; j < t.NumMethod(); j++ {
		fmt.Println("Name", t.Method(j).Name, "Type", t.Method(j).Type)
	}

	v := reflect.ValueOf(&p)
	v.Elem().FieldByName("Name").SetString("chenhongyuan")
	fmt.Printf("%v", p)

}
