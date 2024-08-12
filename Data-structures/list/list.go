package main

import (
	"fmt"
	"container/list"
)

func main() {
	 myList := list.New()
	
	myList.PushBack(1)
	myList.PushBack(12)
	myList.PushBack(15)
	myList.PushBack(333)

	
	for element := myList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
