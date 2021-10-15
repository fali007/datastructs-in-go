package heap

import(
	"fmt"
)

type Node struct{
	Data interface{}
	LeftNode *Node
	RightNode *Node
}

func Heap(a interface{}) *Node{
	return &Node{a,nil,nil}
}

func (n *Node) Add(a interface{}){
	temp:=n
	q:=make([]*Node,0)
	q=append(q,temp)
	for len(q)>0{
		val:=q[0]
		if val.LeftNode!=nil{
			q=append(q,val.LeftNode)
		}else{
			val.LeftNode=&Node{a,nil,nil}
			// fmt.Println("Added as LeftNode")
			break
		}
		if val.RightNode!=nil{
			q=append(q,val.RightNode)
		}else{
			val.RightNode=&Node{a,nil,nil}
			// fmt.Println("Added as RightNode")
			break
		}
		q=q[1:]
	}
}

func (n *Node) minHeapify(){
	if n.LeftNode!=nil{
		n.LeftNode.minHeapify()
		if int(n.Data.(int))>int(n.LeftNode.Data.(int)){
			temp:=n.Data
			n.Data=n.LeftNode.Data
			n.LeftNode.Data=temp
		}
	}
	if n.RightNode!=nil{
		n.RightNode.minHeapify()
		if int(n.Data.(int))>int(n.RightNode.Data.(int)){
			temp:=n.Data
			n.Data=n.RightNode.Data
			n.RightNode.Data=temp
		}
	}
}

func (n *Node) maxHeapify(){
	if n.LeftNode!=nil{
		n.LeftNode.maxHeapify()
		if int(n.Data.(int))<int(n.LeftNode.Data.(int)){
			temp:=n.Data
			n.Data=n.LeftNode.Data
			n.LeftNode.Data=temp
		}
	}
	if n.RightNode!=nil{
		n.RightNode.maxHeapify()
		if int(n.Data.(int))<int(n.RightNode.Data.(int)){
			temp:=n.Data
			n.Data=n.RightNode.Data
			n.RightNode.Data=temp
		}
	}
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

func (n *Node) Print(){
	fmt.Printf("%v - ",n.Data)
	if n.LeftNode!=nil{
		n.LeftNode.Print()
	}
	if n.RightNode!=nil{
		n.RightNode.Print()
	}
}

func (n *Node) PrintIn(){
	if n.LeftNode!=nil{
		n.LeftNode.Print()
	}
	fmt.Printf("%v - ",n.Data)
	if n.RightNode!=nil{
		n.RightNode.Print()
	}
}

func (n *Node) PrintPost(){
	if n.RightNode!=nil{
		n.RightNode.Print()
	}
	fmt.Printf("%v - ",n.Data)
	if n.LeftNode!=nil{
		n.LeftNode.Print()
	}
}

func (n *Node)MaxHeap(){
	fmt.Println("\n performing MaxHeap")
	d:=n.depth()
	for i:=0;i<d;i++{
		n.maxHeapify()
	}
}

func (n *Node)MinHeap(){
	fmt.Println("\n performing MinHeap")
	d:=n.depth()
	for i:=0;i<d;i++{
		n.minHeapify()
	}
}

func (n *Node)PopMin() interface{}{
	n.MinHeap()
	temp:=n.Data
	q:=make([]*Node,0)
	q=append(q,n)
	parent:=make(map[*Node]*Node)
	for len(q)>0{
		val:=q[0]
		if val.LeftNode!=nil{
			q=append(q,val.LeftNode)
			parent[val.LeftNode]=val
		}
		if val.RightNode!=nil{
			q=append(q,val.RightNode)
			parent[val.RightNode]=val
		}
		if val.LeftNode==nil && val.RightNode==nil && len(q)==1{
			n.Data=val.Data
			if parent[val].LeftNode.Data==val.Data{
				parent[val].LeftNode=nil
				return temp
			}
			if parent[val].RightNode.Data==val.Data{
				parent[val].RightNode=nil
				return temp
			}
		}
		q=q[1:]
	}
	return temp
}

func (n *Node)PopMax() interface{}{
	n.MaxHeap()
	temp:=n.Data
	q:=make([]*Node,0)
	q=append(q,n)
	parent:=make(map[*Node]*Node)
	for len(q)>0{
		val:=q[0]
		if val.LeftNode!=nil{
			q=append(q,val.LeftNode)
			parent[val.LeftNode]=val
		}
		if val.RightNode!=nil{
			q=append(q,val.RightNode)
			parent[val.RightNode]=val
		}
		if val.LeftNode==nil && val.RightNode==nil && len(q)==1{
			n.Data=val.Data
			if parent[val].LeftNode.Data==val.Data{
				parent[val].LeftNode=nil
				return temp
			}
			if parent[val].RightNode.Data==val.Data{
				parent[val].RightNode=nil
				return temp
			}
		}
		q=q[1:]
	}
	return temp
}