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
	fmt.Printf("%v - ",n.Data)
	if n.LeftNode!=nil{
		n.LeftNode.Print()
	}
	if n.RightNode!=nil{
		n.RightNode.Print()
	}
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
		fmt.Println("Inserted",t,"as LeftNode of - ", temp.Data)
	}else{
		temp.RightNode=&Node{t,nil,nil}
		fmt.Println("Inserted",t,"as RightNode of - ", temp.Data)
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