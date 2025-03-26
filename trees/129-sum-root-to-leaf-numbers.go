/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	sum := 0
	calculateSum(root, "", &sum)
	return sum
}

func calculateSum(node *TreeNode, currStr string, sum *int) {
    if node == nil {
        return
    }

    currStr += strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		num, _ := strconv.Atoi(currStr)
		*sum += num
		return
	}

	calculateSum(node.Left, currStr, sum)
	calculateSum(node.Right, currStr, sum)
}