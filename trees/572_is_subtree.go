/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var tree string
    serialize(root, &tree)

    var subTree string
    serialize(subRoot, &subTree)

    return strings.Contains(tree, subTree)
}

func serialize(node *TreeNode, str *string) {
    if node == nil {
        *str += ",#"
        return
    }

    *str += "," + strconv.Itoa(node.Val)

    serialize(node.Left, str)
    serialize(node.Right, str)
}
