//  Accessing the Last Element

package main 
 
import(
	"fmt"
	"container/list"
)

func main () {

myList := list.New()

myList.PushBack(1)
myList.PushBack(2)
myList.PushBack(3)
myList.PushBack(4)

lastElement := myList.Back()
if lastElement != nil {
	fmt.Println("Last Element:", lastElement.Value) 
}

	for element := myList.Back(); element != nil; element=element.Prev() {
		fmt.Println(element.Value)
	} 
}