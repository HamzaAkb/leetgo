/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    queue := []*TreeNode{root}
    result := []int{}

    for len(queue) > 0 {
        size, largest := len(queue), queue[0].Val
        for i:=0; i<size; i++ {
            top := queue[i]

            if top.Val > largest {
                largest = top.Val
            }

            if top.Left != nil {
                queue = append(queue, top.Left)
            }

            if top.Right != nil {
                queue = append(queue, top.Right)
            }
        }
        queue = queue[size:]
        result = append(result, largest)
    }

    return result
}