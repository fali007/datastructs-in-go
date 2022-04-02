package main


import (
	"datastructs.com/heap"
	"datastructs.com/tree"
	"datastructs.com/graph"
	"fmt"
)

func main(){
	h:=heap.Heap(0)
	fmt.Println(h)

	t:=tree.BinaryTree(4)
	t.PrintIn()

	v:=map[int][]int{1:{2,3,4},2:{3,5}}
	g:=graph.GenerateGraphAdjacencyList(v)
	fmt.Printf("\n%+v\n",g)
	var a interface{}=10
	fmt.Println(g[0].Bfs(a))
}