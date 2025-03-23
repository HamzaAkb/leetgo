/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        root = &TreeNode{Val: val}
        return root
    }

    if val <= root.Val {
        if root.Left == nil {
            root.Left = &TreeNode{Val: val}
            return root
        }
        insertIntoBST(root.Left, val)
    } else {
        if root.Right == nil {
            root.Right = &TreeNode{Val: val}
            return root
        }
        insertIntoBST(root.Right, val)
    }
    return root
}