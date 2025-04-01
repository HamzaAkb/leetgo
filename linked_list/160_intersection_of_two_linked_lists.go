/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    lengthA, lengthB := 0, 0
    nodeA, nodeB := headA, headB

    for nodeA != nil {
        nodeA = nodeA.Next
        lengthA++
    }

    for nodeB != nil {
        nodeB = nodeB.Next
        lengthB++
    }

    if lengthA > lengthB {
        for i:=0; i < lengthA - lengthB; i++ {
            headA = headA.Next
        }
    } else if lengthB > lengthA {
        for i:=0; i < lengthB - lengthA; i++ {
            headB = headB.Next
        }
    }

    for headA != nil && headB != nil {
        if headA == headB {
            return headA
        }
        headA = headA.Next
        headB = headB.Next
    }

    return nil
}