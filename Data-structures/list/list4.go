// Inserting Elements After a Specific Element

package main

import (
	"fmt"
	"container/list"
)

func main() {
	myList := list.New()

   first := myList.PushBack("First")
   second := myList.PushBack("Second")
   third := myList.PushBack("Third")

  
   myList.InsertAfter("New Element", second)
   
   fmt.Println("First element:", first.Value)
   fmt.Println("Third element:", third.Value)
 
   
   for element := myList.Front(); element !=nil; element = element.Next() {
	fmt.Println(element.Value)
   }
}
