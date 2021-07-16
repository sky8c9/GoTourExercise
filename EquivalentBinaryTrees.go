package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	inorder(t, ch)
	close(ch)
}

func inorder(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	inorder(t.Left, ch)
	ch <- t.Value
	inorder(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	for v:= range ch1 {	
		if v !=  <- ch2 {	
			return false
		}
	}
	
	return true
}

func main() {
	check1 := Same(tree.New(3), tree.New(3))
	fmt.Println(check1)
}