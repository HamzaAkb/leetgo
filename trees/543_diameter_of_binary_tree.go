/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
	calcHeight(root, &maxDiameter)

	return maxDiameter
}

func calcHeight(root *TreeNode, maxDiameter *int) int {
    if root == nil {
        return 0
    }

    leftHeight := calcHeight(root.Left, maxDiameter)
    rightHeight := calcHeight(root.Right, maxDiameter)

    *maxDiameter = int(math.Max(float64(*maxDiameter), float64(leftHeight + rightHeight)))

    return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}