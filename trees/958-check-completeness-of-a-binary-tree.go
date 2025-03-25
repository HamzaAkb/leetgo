/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	nilFound := false

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == nil {
			nilFound = true
		} else if nilFound {
			return false
		} else {
			queue = append(queue, current.Left)
			queue = append(queue, current.Right)
		}
	}

	return true
}