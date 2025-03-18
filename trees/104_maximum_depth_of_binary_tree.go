/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftHeight := 1 + maxDepth(root.Left)
    rightHeight := 1 + maxDepth(root.Right)

    return max(leftHeight, rightHeight)    
}