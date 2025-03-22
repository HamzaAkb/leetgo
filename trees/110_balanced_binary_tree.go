/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    _, result := checkBalance(root)
    return result
}

func checkBalance(root *TreeNode) (int, bool){
    if root == nil {
        return 0, true
    }

    leftHeight, leftBalanced := checkBalance(root.Left)
    rightHeight, rightBalanced := checkBalance(root.Right)

    if !leftBalanced || !rightBalanced {
        return -1, false
    }

    if int(math.Abs(float64(leftHeight) - float64(rightHeight))) > 1 {
        return -1, false
    }

    if leftHeight > rightHeight {
        return 1 + leftHeight, true
    }

    return 1 + rightHeight, true
}