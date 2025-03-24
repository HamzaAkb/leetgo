/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    return traverse(root, targetSum, 0)
}

func traverse(root *TreeNode, targetSum int, currentSum int) bool {
    if root == nil {
        return false
    }

    currentSum += root.Val

    if root.Left == nil && root.Right == nil {
        return currentSum == targetSum
    }

    return traverse(root.Left, targetSum, currentSum) || traverse(root.Right, targetSum, currentSum)
}
