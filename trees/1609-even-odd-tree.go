/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		size := len(queue)
        prevVal := queue[0].Val

		for i := 0; i < size; i++ {
			top := queue[i]

			if level % 2 == 0  {
                if top.Val % 2 == 0 || (i > 0 && top.Val <= prevVal) {
                    return false
                }
			}
            
            if level % 2 != 0 {
                if top.Val % 2 != 0 || (i > 0 && top.Val >= prevVal) {
                    return false
                }
            }

            prevVal = top.Val

			if top.Left != nil {
				queue = append(queue, top.Left)
			}

			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
        
		queue = queue[size:]
		level++
	}

	return true
}