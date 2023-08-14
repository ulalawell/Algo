package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	data []*TreeNode
}

func (queue *Queue) push(node *TreeNode) {
	newData := []*TreeNode{node}
	for _, element := range queue.data {
		newData = append(newData, element)
	}
	queue.data = newData
}

func (queue *Queue) pop() *TreeNode {
	if len(queue.data) != 0 {
		element := queue.data[len(queue.data)-1]
		queue.data = queue.data[:len(queue.data)-1]
		return element
	}
	return nil
}

func (queue *Queue) size() int {
	return len(queue.data)
}

func main() {
	var NodeInput = TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   3,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   13,
				Right: nil,
				Left:  nil,
			},
		},
		Left: &TreeNode{
			Val: 12,
			Right: &TreeNode{
				Val:   -3,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   1488,
				Right: nil,
				Left:  nil,
			},
		},
	}

	var queue = Queue{
		data: []*TreeNode{},
	}
	order := []int{}

	queue.push(&NodeInput)
	order = append(order, NodeInput.Val)

	for queue.size() > 0 {
		node := queue.pop()

		if node != nil {
			if node.Left != nil {
				queue.push(node.Left)
				order = append(order, node.Left.Val)
			}

			if node.Right != nil {
				queue.push(node.Right)
				order = append(order, node.Right.Val)
			}
		}
	}

	fmt.Println(order)
}
