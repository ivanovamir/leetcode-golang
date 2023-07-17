package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// O(n+m) | O(n+m)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)

	return list2
}

// O(n+m) | O(n+m)
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil && list2 == nil {
//		return nil
//	}
//
//	ptr1 := list1
//
//	ptr2 := list2
//
//	for ptr1 != nil {
//		// in every loop iteration move node ptr
//		if ptr2 != nil {
//			if ptr2.Val <= ptr1.Val {
//				if ptr2.Next != nil {
//					if ptr2.Next.Val < ptr1.Val {
//						ptr2 = ptr2.Next
//						continue
//					}
//				}
//				newNode := &ListNode{
//					Val:  ptr1.Val,
//					Next: ptr2.Next,
//				}
//				ptr2.Next = newNode
//				ptr2 = ptr2.Next // next ptr of list2
//				ptr1 = ptr1.Next
//			} else {
//				newNode := &ListNode{
//					Val:  ptr1.Val,
//					Next: ptr2,
//				}
//				list2 = newNode
//				ptr2 = list2
//				ptr1 = ptr1.Next
//				continue
//			}
//		} else {
//			newNode := &ListNode{
//				Val: ptr1.Val,
//			}
//			ptr2 = newNode
//			list2 = newNode
//			ptr1 = ptr1.Next
//		}
//	}
//	return list2
//}

func main() {}
