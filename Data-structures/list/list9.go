// Clearing a List

package main

import(
	"fmt"
	"container/list"
)

func main() {
	myList := list.New()
	
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	
	for myList.Len() > 0 {
		myList.Remove(myList.Front())
	}

	if myList.Len() == 0 {
		fmt.Println("The list has been cleared.")
	}
}
	
