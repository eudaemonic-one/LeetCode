/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Approach #2: Iterative Approach

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil {
        return nil
    }
    var prev, curr *ListNode
    curr = head
    for m > 1 {
        prev = curr
        curr = curr.Next
        m--
        n--
    }
    last, tail := prev, curr
    for n > 0 {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
        n--
    }
    if last != nil {
        last.Next = prev
    } else {
        head = prev
    }
    tail.Next = curr
    return head
}

// Approach #1: Recursive Approach

// var successor *ListNode

// func reverseBetween(head *ListNode, m int, n int) *ListNode {
//     if m == 1 {
//         return reverseN(head, n)
//     }
//     head.Next = reverseBetween(head.Next, m-1, n-1)
//     return head
// }

// func reverseN(head *ListNode, n int) *ListNode {
//     if n == 1 {
//         successor = head.Next
//         return head
//     }
//     last := reverseN(head.Next, n-1)
//     head.Next.Next = head
//     head.Next = successor
//     return last
// }
