/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{0, nil}
    prev := dummy
    for curr := head; curr != nil; {
        next := curr.Next
        for prev.Next != nil && prev.Next.Val < curr.Val {
            prev = prev.Next
        }
        curr.Next = prev.Next
        prev.Next = curr
        prev = dummy
        curr = next
    }
    return dummy.Next
}
