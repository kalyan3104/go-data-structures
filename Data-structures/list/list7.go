//Working with Empty Lists

package main

import (
	"fmt"
	"container/list"
)

func main() {

	myList := list.New()
	if myList.Len() == 0 {
		fmt.Println("the list is empty")
	}
 
	myList.PushBack(43)
	if myList.Len() > 0 {
		fmt.Println("the is list is no longer empty")
	}
}