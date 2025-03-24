/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    var values []int
	inorderTraversal(root, &values)
	
	return calculateMinDiff(values)
}

func inorderTraversal(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	
	inorderTraversal(root.Left, values)
	*values = append(*values, root.Val)
	inorderTraversal(root.Right, values)
}

func calculateMinDiff(values []int) int {
    minDiff := math.MaxInt32
    
    for i:=1; i<len(values); i++ {
        diff := values[i] - values[i-1]
		if  diff < minDiff {
			minDiff = diff
		}
	}

    return minDiff
}