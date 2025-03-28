/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    res := traverse(root)

    if res == -1 {
        return false
    }

    return true
}

func traverse(root *TreeNode) int {
    if root == nil {
        return 0
    }

    lh := traverse(root.Left)
    rh := traverse(root.Right)

    if math.Abs(float64(lh) - float64(rh)) > 1 {
        return -1
    }

    if lh == -1 || rh == -1 {
        return -1
    }

    return 1 + max(lh, rh)
}

func max (a, b int) int {
    if a > b {
        return a
    }
    return b
}