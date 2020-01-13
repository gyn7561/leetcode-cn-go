package _100

//https://leetcode-cn.com/problems/same-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if q == nil && p == nil {
		return true
	}
	if q == nil || p == nil {
		return false
	}
	if q.Val != p.Val {
		return false
	}
	if isSameTree(p.Left, q.Left) {
		return isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}
