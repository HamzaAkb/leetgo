/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var result []int
    traverse(root, &result)
    return result
}

func traverse(root *TreeNode, result *[]int){
    if root == nil {
        return
    }

    *result = append(*result, root.Val)
    traverse(root.Left, result)
    traverse(root.Right, result)
}