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
	t.PrintIn()

}