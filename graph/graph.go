package graph


type Node struct{
	Data interface{}
	Neighbours []*Node
}

func graph(data interface{})*Node{
	root:=Node{data,make([]*Node,0)}
	return &root
}

func (g *Node)getNode(f interface{})*Node{
	v:=make([]*Node,0)
	v=append(v,g)
	for len(v)>0{
		temp:=v[0]
		if temp.Data==f{
			return temp
		}else{
			for _,i:=range temp.Neighbours{
				v=append(v,i)
			}
		}
		v=v[1:]
	}
	return nil
}

func GenerateGraphAdjacencyList(m map[int][]int)[]*Node{
	g:=graph(0)
	vertices:=make([]*Node, 0)
	for k,v:=range m{
		node:=g.getNode(k)
		if node==nil{
			node=graph(k)
			vertices=append(vertices,node)
		}
		for i:=range v{
			temp:=g.getNode(i)
			if temp!=nil{
				node.Neighbours=append(node.Neighbours,temp)
			}else{
				newNod:=graph(i)
				node.Neighbours=append(node.Neighbours,newNod)
				vertices=append(vertices,newNod)
			}
		}
	}
	return vertices
}

func GenerateGraphMatrix(m []int[]int)[]*Node{
	
}

func (n *Node)Bfs(val interface{})*Node{
	q:=make([]*Node,0)
	q=append(q,n)
	for len(q)>0{
		temp:=q[0]
		q=q[1:]
		if temp.Data==val{
			return temp
		}else{
			for _,i := range temp.Neighbours{
				q=append(q,i)
			}
		}
	}
	return nil
}