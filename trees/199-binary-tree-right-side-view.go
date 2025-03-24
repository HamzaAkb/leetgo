/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    var result []int
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue)

        for i:=0; i<size; i++ {
            top := queue[0]

            if i == size - 1 {
                result = append(result, top.Val)
            }

            queue = queue[1:]
            
            if top.Left != nil {
                queue = append(queue, top.Left)
            }

            if top.Right != nil {
                queue = append(queue, top.Right)
            }
        }
    }

    return result
}