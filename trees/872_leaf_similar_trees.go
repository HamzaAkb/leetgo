/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var leafSequenceRoot1 []int
    var leafSequenceRoot2 []int

    calculateLeafSequence(root1, &leafSequenceRoot1)
    calculateLeafSequence(root2, &leafSequenceRoot2)

    if len(leafSequenceRoot1) != len(leafSequenceRoot2){
        return false
    }

    for i:=0; i<len(leafSequenceRoot1); i++ {
        if leafSequenceRoot1[i] != leafSequenceRoot2[i]{
            return false
        }
    }

    return true
}

func calculateLeafSequence(root* TreeNode, sequence *[]int){
    if root == nil {
        return
    }

    if root.Right == nil && root.Left == nil {
        *sequence = append(*sequence, root.Val)
    } else {
        calculateLeafSequence(root.Left, sequence)
        calculateLeafSequence(root.Right, sequence)
    }
}