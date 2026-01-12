package main

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res := dfs(root)
	return res.node
}

func dfs(node *TreeNode) result {
	if node == nil {
		return result{nil, 0}
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	if left.depth > right.depth {
		return result{left.node, left.depth + 1}
	} else if right.depth > left.depth {
		return result{right.node, right.depth + 1}
	} else {
		return result{node, left.depth + 1}
	}
}

// func main() {
// 	// Example usage:
// 	// Construct a binary tree and call subtreeWithAllDeepest
// 	root := &TreeNode{Val: 3}
// 	root.Left = &TreeNode{Val: 5}
// 	root.Right = &TreeNode{Val: 1}
// 	root.Left.Left = &TreeNode{Val: 6}
// 	root.Left.Right = &TreeNode{Val: 2}

// 	print(subtreeWithAllDeepest(root).Val)
// }

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type result struct {
	node  *TreeNode
	depth int
}
