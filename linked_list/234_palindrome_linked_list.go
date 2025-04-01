/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var vals []int
    node := head

    for node != nil {
        vals = append(vals, node.Val)
        node = node.Next
    }

    length := len(vals)

    for i:=0; i<length; i++ {
        if head.Val != vals[length - i - 1] {
            return false
        }
        head = head.Next
    }

    return true
}