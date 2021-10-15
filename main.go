package main

import(
	"fmt"
	"datastructs.com/heap"
)

func main(){
	val:=heap.Heap(0)

	for i:=1;i<10;i++{
		val.Add(i)
	}

	fmt.Println("\nPre order print")
	val.Print()
	fmt.Println("\nIn order print")
	val.PrintIn()
	fmt.Println("\nPost order print")
	val.PrintPost()

	val.MinHeap()
	fmt.Println("\nPre order print")
	val.Print()

	val.MaxHeap()
	fmt.Println("\nPre order print")
	val.Print()

	fmt.Println("\n - pop min value :",val.PopMin())
	fmt.Println("\nPre order print")
	val.Print()

	fmt.Println("\n - pop max value :",val.PopMax())
	fmt.Println("\nPre order print")
	val.Print()
}