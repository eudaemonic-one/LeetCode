/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{-1, l1}
    curr := dummy
    sum, carry, remainder := 0, 0, 0
    l, r := 0, 0
    for l1 != nil || l2 != nil {
        if l1 != nil {
            l = l1.Val
        } else {
            l = 0
        }
        if l2 != nil {
            r = l2.Val
        } else {
            r = 0
        }
        sum = l + r + carry
        carry, remainder = sum / 10, sum % 10
        curr.Next = &ListNode{remainder, nil}
        curr = curr.Next
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    if carry > 0 {
        curr.Next = &ListNode{carry, nil}
    }
    return dummy.Next
}
