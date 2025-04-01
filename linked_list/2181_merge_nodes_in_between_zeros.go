/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    sum, curr := 0, &ListNode{Next: nil}
    res := curr

    for head != nil {
        if head.Val == 0 && sum != 0 {
            node := &ListNode{Val: sum}
            curr.Next = node
            curr = curr.Next
            sum = 0
        }
        sum += head.Val
        head = head.Next
    }

    return res.Next
}