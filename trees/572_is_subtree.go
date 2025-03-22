/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var queue []*TreeNode
    queue = append(queue, root)

    for len(queue) > 0 {
        top := queue[0]
        queue = queue[1:]

        if checkSubtree(top, subRoot) {
            return true
        }

        if top.Left != nil {
            queue = append(queue, top.Left)
        }

        if top.Right != nil {
            queue = append(queue, top.Right)
        }
    }

    return false
}

func checkSubtree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil || q == nil {
        return false
    }

    if p.Val != q.Val {
        return false
    }

    leftTree := checkSubtree(p.Left, q.Left)
    rightTree := checkSubtree(p.Right, q.Right)

    if !leftTree || !rightTree {
        return false
    }

    return true
}