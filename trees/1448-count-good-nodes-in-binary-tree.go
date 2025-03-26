/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    count := 0
    countGoodNodes(root, &count, root.Val)
    return count
}

func countGoodNodes(node *TreeNode, count *int, largest int){
    if node == nil {
        return
    }

    if node.Val >= largest {
        *count++
        largest = node.Val
    }

    countGoodNodes(node.Left, count, largest)
    countGoodNodes(node.Right, count, largest)
}