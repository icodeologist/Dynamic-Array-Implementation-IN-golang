package main

import (
	"fmt"
)

func main() {
	d := DynamicArray{}
	d.NewDArray()
	for i := range 1000 {
		d.Push(i)
	}
	fmt.Printf("D.data : %v\n", d.data)
	fmt.Println(d.size)
	fmt.Println(len(d.data))
}
