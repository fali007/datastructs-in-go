package tree

import (
	"fmt"
)

type Node struct{
	Data interface{}
	LeftNode *Node
	RightNode *Node
}

func BinaryTree(t interface{})*Node{
	return &Node{t, nil, nil}
}

func (n *Node) Print(){
	if n==nil{
		return
	}
	fmt.Printf("%v - ",n.Data)
	n.LeftNode.Print()
	n.RightNode.Print()
}

func (n *Node) PrintIn(){
	if n==nil{
		return
	}
	n.LeftNode.PrintIn()
	fmt.Printf("%v - ",n.Data)
	n.RightNode.PrintIn()
}

func (n *Node)Add(t interface{}){
	var temp *Node
	for n!=nil{
		temp=n
		if int(t.(int))<int(n.Data.(int)){
			n=n.LeftNode
		}else{
			n=n.RightNode
		}
	}
	if int(t.(int))<int(temp.Data.(int)){
		temp.LeftNode=&Node{t,nil,nil}
		// fmt.Println("Inserted",t,"as LeftNode of - ", temp.Data)
	}else{
		temp.RightNode=&Node{t,nil,nil}
		// fmt.Println("Inserted",t,"as RightNode of - ", temp.Data)
	}
}

func (n *Node)Search(t interface{}) *Node{
	if int(n.Data.(int))==int(t.(int)){
		return n
	}
	if int(n.Data.(int))>int(t.(int)){
		return n.LeftNode.Search(t)
	}else{
		return n.RightNode.Search(t)
	}
}

func (n *Node) getInorder(i []interface{})[]interface{}{
	if n==nil{
		return i
	}
	i=n.LeftNode.getInorder(i)
	i=append(i,n.Data)
	i=n.RightNode.getInorder(i)
	return i
}

func createTree(val []interface{}, s, e int)*Node{
	if s>e{
		return nil
	}
	mid:=(s+e)/2
	root:=Node{val[mid],nil,nil}
	root.LeftNode=createTree(val,s,mid-1)
	root.RightNode=createTree(val,mid+1,e)
	return &root
}

func (n *Node) BalanceTree()*Node{
	var j []interface{}

	fmt.Println("Depth of input tree -",n.depth())

	j=n.getInorder(j)
	n=createTree(j,0,len(j)-1)

	fmt.Println("Depth after balancing - ", n.depth())
	return n
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func (n *Node) depth()int{
	if n==nil {
		return 0
	}
	return 1+max(n.LeftNode.depth(),n.RightNode.depth())
}