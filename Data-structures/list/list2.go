package main

import (
	"fmt"
	"container/list"
)

func main(){
	var intList list.List
	intList.PushBack(1111)
	intList.PushBack(333)
	intList.PushBack(808)
	intList.PushBack(909)

	for element := intList.Front(); element !=nil; element=element.Next(){
		fmt.Println(element.Vakue.(int))
	}

}