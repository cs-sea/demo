package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	k  int
	pv int
}

var v1, v2 Node

func isCousins(root *TreeNode, x int, y int) bool {
	dfs(root, x, y, 0)
	fmt.Println(v1, v2)
	return true
}

func dfs(n *TreeNode, x, y, k int) {
	dfs(n.Left, x, y, k+1)
	dfs(n.Right, x, y, k+1)
	if n.Left != nil && n.Left.Val == x {
		v1.pv = n.Val
		v1.k = k + 1
	}

	if n.Right != nil && n.Right.Val == x {
		v1.pv = n.Val
		v1.k = k + 1
	}
	if n.Left != nil && n.Left.Val == y {
		v2.pv = n.Val
		v2.k = k + 1
	}

	if n.Right != nil && n.Right.Val == y {
		v2.pv = n.Val
		v2.k = k + 1
	}
}
