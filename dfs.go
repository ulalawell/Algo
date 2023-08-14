package main

import "fmt"

type Node struct {
	Val int
	R   *Node
	L   *Node
}

func dfs(node *Node) int {
	if node == nil {
		return 0
	}

	return node.Val + dfs(node.L) + dfs(node.R)
}

func main() {
	var NodeInput = Node{
		Val: 1,
		R: &Node{
			Val: 2,
			R: &Node{
				Val: 3,
				R:   nil,
				L:   nil,
			},
			L: &Node{
				Val: 3,
				R:   nil,
				L:   nil,
			},
		},
		L: nil,
	}

	fmt.Print(dfs(&NodeInput))

}
