/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    slow, fast := head, head
    for fast != nil && fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    second := slow.Next
    slow.Next = nil
    
    l := sortList(head)
    r := sortList(second)
    return merge(l, r)
}

func merge(l, r *ListNode) *ListNode {
    if l == nil {
        return r
    } else if r == nil {
        return l
    }
    
    if l.Val > r.Val {
        l, r = r, l
    }
    head, curr := l, l
    l = l.Next
    
    for l != nil && r != nil {
        if l.Val < r.Val {
            curr.Next = l
            l = l.Next
        } else {
            curr.Next = r
            r = r.Next
        }
        curr = curr.Next
    }
    if l == nil {
        curr.Next = r
    } else if r == nil {
        curr.Next = l
    }
    
    return head
}
