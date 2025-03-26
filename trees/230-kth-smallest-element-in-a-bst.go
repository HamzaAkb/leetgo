/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    nums := []int{}
    traverse(root, &nums, k)
    return nums[k - 1]
}

func traverse(node *TreeNode, nums *[]int, k int){
    if node == nil {
        return
    }

    if len(*nums) > k {
        return
    }

    traverse(node.Left, nums, k)
    *nums = append(*nums, node.Val)
    traverse(node.Right, nums, k)
}