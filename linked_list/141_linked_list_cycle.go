/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)

    for head != nil {
        if seen[head] {
            return true
        }

        seen[head] = true
        head = head.Next
    }

    return false
}