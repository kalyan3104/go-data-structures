package main

import (
	"fmt"
	"container/list"
)
func main() {
	myList := list.New()

	myList.PushBack("programming")
	myList.PushFront("go")
	myList.PushBack("Language")
	
	for element := myList.Front(); element !=nil; element=element.Next(){
		fmt.Println(element.Value)
	 }
}