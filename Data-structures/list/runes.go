package main

import (
	"container/list"
	"fmt"
)

func main() {

	key := list.New()

	name := ""
	for _, char := range name {
		key.PushBack(char)
	}

	for element := key.Back(); element != nil; element = element.Prev() {
		fmt.Print(string(element.Value.(rune)))
	}
}
  
