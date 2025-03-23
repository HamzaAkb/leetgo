/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }

    result := strconv.Itoa(root.Val)

    if root.Left != nil {
        result = fmt.Sprintf("%s(%s)", result, tree2str(root.Left))
    } else if  root.Right != nil {
        result = fmt.Sprintf("%s()", result)
    }

    if root.Right != nil {
        result = fmt.Sprintf("%s(%s)", result, tree2str(root.Right))
    }

    return result
}