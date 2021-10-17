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
	t.Add(3)
	t.Add(2)
	t.Add(1)
	t.Add(5)
	t.Add(6)
	t.Add(7)
	t.Add(8)
	t.Add(9)
	t.Add(10)
	t.Add(11)
	t.Add(12)
	t.Add(13)
	t.Add(14)
	t.Add(15)
	fmt.Println("\nInorder")
	t.PrintIn()

	var i interface{} =2
	fmt.Println("\nSearch Result :",t.Search(i))

	t=t.BalanceTree()
	t.Print()
}