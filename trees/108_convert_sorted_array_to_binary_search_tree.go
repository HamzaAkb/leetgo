/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return buildBST(nums, 0, len(nums) - 1)
}

func buildBST(nums []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }

    mid := (left + right) / 2
    node := &TreeNode{Val: nums[mid]}
    node.Left = buildBST(nums, left, mid-1)
    node.Right = buildBST(nums, mid+1, right)

    return node
}