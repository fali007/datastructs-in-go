package main

import(
	"datastructs.com/heap"
	"testing"
)

func TestHeap(t *testing.T){
	val:=heap.Heap(0)

	for i:=1;i<10;i++{
		val.Add(i)
	}

	got:=val.PopMin()
	want:=0

	if got!=want{
		t.Errorf("Got %q , want %q \n", got, want)
	}

	got=val.PopMax()
	want=9

	if got!=want{
		t.Errorf("Got %q , want %q \n", got, want)
	}
}