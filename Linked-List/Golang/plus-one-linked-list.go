/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    dummy := &ListNode{0, head}
    notNine := dummy
    for head != nil {
        if head.Val != 9 {
            notNine = head
        }
        head = head.Next
    }
    notNine.Val++
    nine := notNine.Next
    for nine != nil {
        nine.Val = 0
        nine = nine.Next
    }
    if dummy.Val == 0 {
        return dummy.Next
    }
    return dummy
}
