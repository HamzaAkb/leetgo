/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    queue := []*TreeNode{root}
    leftMost := 0

    for len(queue) > 0 {
        size := len(queue)
        for i:=0; i<size; i++ {
            top := queue[i]

            if i == 0 {
                leftMost = top.Val
            }

            if top.Left != nil {
                queue = append(queue, top.Left)
            }

            if top.Right != nil {
                queue = append(queue, top.Right)
            }
        }
        queue = queue[size:]
    }

    return leftMost
}