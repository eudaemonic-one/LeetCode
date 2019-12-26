/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{0,l1}
    prev := dummy
    carry, remainder := 0, 0
    for l1 != nil && l2 != nil {
        sum := l1.Val + l2.Val + carry
        carry, remainder = sum/10, sum%10
        l1.Val = remainder
        prev, l1, l2 = prev.Next, l1.Next, l2.Next
    }
    if l2 != nil && prev != nil {
        prev.Next = l2
    }
    for l1 = prev.Next; l1 != nil && carry > 0; l1 = l1.Next {
        sum := l1.Val + carry
        carry, remainder = sum/10, sum%10
        l1.Val = remainder
        prev = l1
    }
    if carry > 0 {
        prev.Next = &ListNode{carry, nil}
    }
    return dummy.Next
}
