package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	lst := l
	for lst != nil {
		fmt.Println(lst.Val)
		lst = lst.Next
	}
}

// O(n) | O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	newList := &ListNode{}
	current := newList

	p1 := head

	p2 := head

	i := n

	for i != 0 {
		p2 = p2.Next
		i--
	}

	for p2 != nil {
		current.Next = p1      // newList.Next = p1
		current = current.Next // change current address

		p1 = p1.Next

		if p2.Next != nil {
			p2 = p2.Next
		} else {
			break
		}
	}

	current.Next = p1.Next
	current = current.Next
	return newList.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	removeNthFromEnd(head, 1).Print()
}
