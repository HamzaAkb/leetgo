import (
    "slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for i := 0; len(queue) > 0; i++ {
		var items []int
		size := len(queue)

		for j := 0; j < size; j++ {
			top := queue[j]
			items = append(items, top.Val)

			if top.Left != nil {
				queue = append(queue, top.Left)
			}

			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}

        if i % 2 != 0 {
            slices.Reverse(items) 
        }

		result = append(result, items)
		queue = queue[size:]
	}

	return result
}