/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var result []int
    traverse(root, &result)
    return result
}

func traverse(root *TreeNode, result *[]int){
    if root == nil {
        return
    }

    traverse(root.Left, result)
    traverse(root.Right, result)
    *result = append(*result, root.Val)
}