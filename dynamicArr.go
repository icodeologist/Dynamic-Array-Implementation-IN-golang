package main

import (
	"fmt"
)

type DynamicArray struct {
	data     []int
	size     int
	capacity int
}

func (d *DynamicArray) NewDArray() *DynamicArray {
	return &DynamicArray{
		data:     make([]int, 0, 0),
		size:     0,
		capacity: 0,
	}
}

func (d *DynamicArray) GetSize() int {
	return d.size
}

func (d *DynamicArray) GetCapacity() int {
	return d.capacity
}
func (d *DynamicArray) GetData() []int {
	return d.data
}
func (d *DynamicArray) InsertATindex(index int, value int) error {
	if index < 0 || index >= d.size {
		err := fmt.Errorf("Index %d is out of range", index)
		panic(err)
	}
	d.data[index] = value
	fmt.Printf("Array after insertion : %v\n", d.data)
	return nil
}

// ar[i] = x
func (d *DynamicArray) Push(x int) {
	if d.size < d.capacity {
		// Still can push
		d.data[d.size] = x
	} else {
		d.Resize()
		d.data[d.size] = x
	}
	d.size++
}
func (d *DynamicArray) Resize() {
	newCap := d.capacity
	if newCap == 0 {
		newCap = 1
	} else {
		newCap *= 2
	}
	newData := make([]int, newCap)
	for i := 0; i < d.size; i++ {
		newData[i] = d.data[i]
	}
	d.capacity = newCap
	d.data = newData
}
