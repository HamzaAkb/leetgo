/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        } else {
            node := findMin(root.Right)
            root.Val = node.Val
            root.Right = deleteNode(root.Right, node.Val)
        }
    }

    return root
}

func findMin(node *TreeNode) *TreeNode {
    prev := node
    for node != nil {
        prev = node
        node = node.Left
    }
    return prev
}