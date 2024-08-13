package main

import(
	"fmt"
	"container/list"
)

func main() {

	myList := list.New()
	numList := list.New()
	
	k := list.New()
	
	name := "kalyan"
	 for _, char := range name {
		k.PushBack(char)
	 }

	myList.PushFront(1111)
	numList.PushBack(222)
	numList.PushBack(333)
	numList.PushBack(808)
	numList.PushBack(909)
	numList.PushBack(404)

	fmt.Println("Reverse string name of 'k':")
	for element := k.Back(); element !=nil; element = element.Prev() {
		fmt.Println(string(element.Value.(rune)))
	}  
	
    fmt.Println("number of 'myList':")
	for element := myList.Front(); element !=nil; element = element.Next() {
		fmt.Println(element.Value)
	}
 
	fmt.Println("numbers in reverse in 'numList':")
	for element := numList.Back(); element !=nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}