/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Approach #2: Recursive Approach

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    last := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return last
}

// Approach #1: Iterative Approach

// func reverseList(head *ListNode) *ListNode {
//     var prev *ListNode
//     var cur = head
//     for cur != nil {
//         next := cur.Next
//         cur.Next = prev
//         prev = cur
//         cur = next
//     }
//     return prev
// }
