/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return traverse(p, q)
}

func traverse(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if (p == nil || q == nil) || (p.Val != q.Val) {
        return false
    }   

    return traverse(p.Left, q.Left) && traverse(p.Right, q.Right)
}