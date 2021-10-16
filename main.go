package main


import (
	"datastructs.com/heap"
	"datastructs.com/tree"
	"fmt"
)

func main(){
	h:=heap.Heap(0)
	fmt.Println(h)

	t:=tree.BinaryTree(4)
	t.Add(2)
	t.Add(1)
	t.Add(6)
	t.Add(5)
	t.Add(3)
	t.Print()

	var i interface{} =2
	fmt.Println(t.Search(i))
}