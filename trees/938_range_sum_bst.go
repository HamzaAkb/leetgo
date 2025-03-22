/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    var nums []int

    traverse(root, &nums, low, high)

    var sum int

    for _, num := range nums {
        sum += num
    }

    return sum
}

func traverse(root *TreeNode, nums *[]int, low int, high int){
    if root == nil {
        return
    }

    if root.Val >= low && root.Val <= high {
        *nums = append(*nums, root.Val)
    }

    traverse(root.Left, nums, low, high)
    traverse(root.Right, nums, low, high)
}