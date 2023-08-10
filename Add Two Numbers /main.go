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

// O(n+m) | O(n+m)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newList := &ListNode{}

	current := newList

	nextSumPlus := 0

	for l1 != nil || l2 != nil {
		sum := nextSumPlus

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		nextSumPlus = sum / 10

		current = current.Next
	}

	if nextSumPlus > 0 {
		current.Next = &ListNode{
			Val:  nextSumPlus,
			Next: nil,
		}
	}

	return newList.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	addTwoNumbers(l1, l2).Print()
}
