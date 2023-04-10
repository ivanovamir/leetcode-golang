package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//firstNode := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//secondNode := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	//firstNode := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//secondNode := &TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	//firstNode := &TreeNode{}
	//secondNode := &TreeNode{}

	//firstNode := &TreeNode{}
	//secondNode := &TreeNode{
	//	Val:  0,
	//	Left: nil,
	//}

	firstNode := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	secondNode := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isSameTree(firstNode, secondNode))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q != nil || p != nil && q == nil {
		return false
	}

	queue1 := make([]*TreeNode, 0)
	queue1 = append(queue1, p)

	queue2 := make([]*TreeNode, 0)
	queue2 = append(queue2, q)

	for len(queue1) > 0 || len(queue2) > 0 {

		if len(queue2) != len(queue1) {
			return false
		}

		vertex1 := queue1[0]
		vertex2 := queue2[0]

		if vertex1 == nil && vertex2 == nil {
			queue1 = queue1[1:]
			queue2 = queue2[1:]
			continue
		}

		if vertex1 == nil || vertex2 == nil {
			return false
		}

		queue1 = queue1[1:]
		queue2 = queue2[1:]

		if vertex1.Val != vertex2.Val {
			return false
		}

		if vertex1 != nil {
			queue1 = append(queue1, vertex1.Left, vertex1.Right)
		}

		if vertex2 != nil {
			queue2 = append(queue2, vertex2.Left, vertex2.Right)
		}
	}

	return true
}
