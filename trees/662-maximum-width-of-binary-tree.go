/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Node struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	maxWidth, queue := 0, []Node{{root, 0}}

	for len(queue) > 0 {
		size := len(queue)

        width := queue[size - 1].index - queue[0].index + 1

        if width > maxWidth {
            maxWidth = width
        }

		for i := 0; i < size; i++ {
			top := queue[0]
			node, index := top.node, top.index
            queue = queue[1:]
            

			if node.Left != nil {
				queue = append(queue, Node{node: node.Left, index: 2 * index})
			}

			if node.Right != nil {
				queue = append(queue, Node{node: node.Right, index: 2*index + 1})
			}
		}
	}

	return maxWidth
}