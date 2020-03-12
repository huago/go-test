package test

import (
	"encoding/json"
	"fmt"
)

type Nums struct {
	Num     int    `json:"num"`
	Cap     int    `json:"cap"`
	Address string `json:"address"`
	Lux     int    `json:"lux"`
}

func Copy() {
	nums := [3]int{}
	nums[0] = 1

	fmt.Printf("nums:%v, len:%d, cap:%d\n", nums, len(nums), cap(nums))

	dnums := nums[0:1]
	dnums[0] = 5

	fmt.Printf("nums:%v, len:%d, cap:%d\n", nums, len(nums), cap(nums))
	fmt.Printf("dnums:%v, len:%d, cap:%d\n", dnums, len(dnums), cap(dnums))

	s := "{\"num\":1, \"cap\":2, \"address\":\"hello\", \"sex\":1}"

	n := Nums{}
	err := json.Unmarshal([]byte(s), &n)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Printf("Nums=%v\n", n)
}
