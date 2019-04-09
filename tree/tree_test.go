package tree

import (
	"fmt"
	"testing"
)

func TestBst_Insert(t *testing.T) {
	bst := NewBst(5)
	bst.Insert(4)
	bst.Insert(6)
	bst.Insert(3)
	bst.Insert(8)
	bst.InOrderDisplay()
	ok, node := bst.Find(5)
	fmt.Println(ok)
	fmt.Println(node)
	bst.Del(5)
	ok, node1 := bst.Find(5)
	fmt.Println(node == node1)
	fmt.Println(ok)
	fmt.Println(node1)
	bst.InOrderDisplay()

}
