package _2

//https://leetcode-cn.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//func printListNode(l *ListNode) {
//	c := l
//	for {
//		fmt.Print(c.Val)
//		c = c.Next
//		if c == nil {
//			fmt.Println()
//			return
//		}
//	}
//}
//
//
//func reverseListNodeAndGetLength(l *ListNode) (*ListNode, int) {
//	printListNode(l)
//	arr := make([]*ListNode, 1)
//	arr[0] = l
//	current := l
//	for {
//		if current.Next != nil {
//			current = current.Next
//			arr = append(arr, current)
//		} else {
//			break
//		}
//	}
//	for i := range arr {
//		item := arr[i]
//		if i == 0 {
//			item.Next = nil
//		} else if i == len(arr)-1 {
//			item.Next = arr[len(arr)-2]
//		} else {
//			item.Next = arr[i-1]
//		}
//	}
//	printListNode(arr[len(arr)-1])
//	return arr[len(arr)-1], len(arr)
//}

func listNodeLen(l *ListNode) int {
	arr := make([]*ListNode, 1)
	arr[0] = l
	current := l
	for {
		if current.Next != nil {
			current = current.Next
			arr = append(arr, current)
		} else {
			break
		}
	}
	return len(arr)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := listNodeLen(l1)
	len2 := listNodeLen(l2)
	c1 := l1
	c2 := l2
	if len2 > len1 {
		c1 = l2
		c2 = l1
	}
	current := c1
	result := c1
	//十进制进位
	greaterThanEqual10 := false
	for {
		if c1 != nil && c2 != nil {
			value := c2.Val + c1.Val
			if greaterThanEqual10 {
				value++
			}
			greaterThanEqual10 = value >= 10
			c1.Val = value % 10
			current = c1
			c1 = c1.Next
			c2 = c2.Next
		} else if c1 != nil && c2 == nil {
			value := c1.Val
			if greaterThanEqual10 {
				value++
			}
			greaterThanEqual10 = value >= 10
			c1.Val = value % 10
			current = c1
			c1 = c1.Next
		} else if c1 == nil && c2 == nil {
			if greaterThanEqual10 {
				current.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
			}
			break
		}
	}
	return result
}
