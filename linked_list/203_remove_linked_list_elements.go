/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    node := &ListNode{}
    result := node

    for head != nil {
        if head.Val != val {
            node.Next = head
            node = node.Next
        }
        head = head.Next
    }

    node.Next = nil

    return result.Next
}