package main

import (
	"fmt"
	"strconv"

	"github.com/ReanGD/go-algo/hmap"
)

func main() {
	fmt.Println("run")
	m := hmap.New(hmap.HashString)
	for i := 0; i != 50; i++ {
		m.Insert(strconv.Itoa(i), i*10)
	}
	for i := 0; i != 50; i++ {
		v, ok := m.Get(strconv.Itoa(i))
		if !ok {
			fmt.Println("not found")
		} else {
			fmt.Println(v)
		}
	}
	fmt.Println("finish")
}
