package _21

//https://leetcode-cn.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var first *ListNode
	c1 := l1
	c2 := l2

	if l1.Val > l2.Val {
		first = l2
		c2 = first.Next
	} else {
		first = l1
		c1 = l1.Next
	}
	c := first
	for {
		if c1 == nil && c2 == nil {
			break
		}
		if c1 == nil && c2 != nil {
			c.Next = c2
			c2 = c2.Next
		} else if c2 == nil && c1 != nil {
			c.Next = c1
			c1 = c1.Next
		} else if c1.Val < c2.Val {
			c.Next = c1
			c1 = c1.Next
		} else {
			c.Next = c2
			c2 = c2.Next
		}
		c = c.Next
	}
	return first
}
