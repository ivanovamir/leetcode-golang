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

// O(n) | O(n + n) = O(2n) = O(n)
//func deleteDuplicates(head *ListNode) *ListNode {
//	repeat := make(map[int]struct{})
//
//	newList := &ListNode{}
//	current := newList
//
//	for head != nil {
//
//		_, ok := repeat[head.Val]
//
//		if !ok {
//			current.Next = &ListNode{
//				Val:  head.Val,
//				Next: nil,
//			}
//			current = current.Next
//			repeat[head.Val] = struct{}{}
//		}
//
//		head = head.Next
//	}
//
//	return newList.Next
//}

// O(n) | O(n + n) = O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head
	for currentNode != nil && currentNode.Next != nil {
		next := currentNode.Next
		if currentNode.Val == next.Val {
			if next.Next == nil {
				currentNode.Next = nil
			} else {
				currentNode.Next = next.Next
			}
		} else {
			currentNode = currentNode.Next
		}
	}
	return head
}

func main() {
	lst := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	deleteDuplicates(lst).Print()
}
