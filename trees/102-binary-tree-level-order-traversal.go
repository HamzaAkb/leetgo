/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    queue := []*TreeNode{root}


    for len(queue) > 0 {
        var level []int
        size := len(queue)

        for i:=0; i<size; i++ {
            top := queue[0]
            queue = queue[1:]

            level = append(level, top.Val)

            if top.Left != nil {
                queue = append(queue, top.Left)
            }

            if top.Right != nil {
                queue = append(queue, top.Right)
            }
        }

        result = append(result, level)
    }

    return result
}