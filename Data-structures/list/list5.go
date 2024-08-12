//  Removing Elements from the List

package main
  
import (
	"fmt"
	"container/list"
)

func main () {

	myList := list.New()

	myList.PushBack(2)
	myList.PushBack(8)
	myList.PushBack(9)
	myList.PushBack(1)
	myList.PushBack(3)
	myList.PushBack(5)
	myList.PushBack(6)

	elementToRemove := myList.Front().Next()
	myList.Remove(elementToRemove)
	
	for element := myList.Front(); element !=nil; element=element.Next() {
		fmt.Println(element.Value)
	}

}


