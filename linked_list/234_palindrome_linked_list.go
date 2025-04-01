/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    curr := slow
    var prev *ListNode

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    for head != nil && prev != nil {
        if head.Val != prev.Val {
            return false
        }
        head = head.Next
        prev = prev.Next
    }

    return true
}