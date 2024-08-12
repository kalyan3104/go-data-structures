package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()

	
	name := "kalyan"
	for _, char := range name {
		myList.PushBack(char)
	}

	
	for element := myList.Back(); element != nil; element = element.Prev() {
		fmt.Print(string(element.Value.(rune)))
	}
	fmt.Println()
}
