//Implementing Side-by-Side Output Display

package main

import (
	"fmt"
	"container/list"
)


func main() {

	myList := list.New()
	numList := list.New()
	k := list.New()

    myList.PushFront(111)
    numList.PushBack(1111)
	numList.PushBack(333)
	numList.PushBack(808)
	numList.PushBack(909)	
	
	name := "kalyan"
	 for _, char := range name {
		k.PushBack(char)
	 }
	 
	 fmt.Println("reverse string of 'k:")
    for element := k.Back(); element !=nil; element = element.Prev() {
		fmt.Print(string(element.Value.(rune)), " ")
	}
     
	fmt.Println()

	fmt.Println("number of 'myList:")
    for element := myList.Front(); element !=nil; element = element.Next() {
		fmt.Print(element.Value)
	}
    
	fmt.Println()

	 fmt.Println("numbers of 'numList:")
	 for element := numList.Back(); element !=nil; element = element.Prev() {
		fmt.Println(element.Value)
	 }
      fmt.Println()
	
	}