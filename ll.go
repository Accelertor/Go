package main

import (
	"fmt"
)

type Node struct {
	val int
	n   *Node
}

var n = &Node{}

func add(c int) {
	var x = &Node{val: c, n: nil}
	if n.n == nil {
		n.n = x
	}
}
func main() {
	for i := 1; i != 10; i++ {
		add(i)
	}
	for n.n != nil {
		fmt.Println(n.val)
	}
}
