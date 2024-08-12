package main

import (
	"fmt"
	"container/list"
)

func main() {

	myList := list.New()

	myList.PushBack(100)
	myList.PushBack(200)
	myList.PushBack(300)

	for element := myList.Front(); element != nil; element = element.Next() {
		applyFunction(element.Value)
	}
}

func applyFunction(value interface{}) {
	fmt.Println(value)
}
