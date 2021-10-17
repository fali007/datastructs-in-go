package tree

import(
	"testing"
)

func TestBinaryTree(r *testing.T){
	t:=BinaryTree(4)
	t.Add(3)
	t.Add(2)
	t.Add(1)
	for i:=5;i<16;i++{
		t.Add(i)
	}
	

	got:=t.depth()
	want:=12

	if got!=want{
		r.Errorf("Got %q , want %q \n", got, want)
	}

	t=t.BalanceTree()

	got=t.depth()
	want=4

	if got!=want{
		r.Errorf("Got %q , want %q \n", got, want)
	}

	var i interface{} =2
	got=int(t.Search(i).Data.(int))
	want=2

	if got!=want{
		r.Errorf("Got %q , want %q \n", got, want)
	}
}