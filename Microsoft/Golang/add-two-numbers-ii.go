/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    st1 := make([]int, 0)
    st2 := make([]int, 0)
    for ; l1 != nil; l1 = l1.Next {
        st1 = append(st1, l1.Val)
    }
    for ; l2 != nil; l2 = l2.Next {
        st2 = append(st2, l2.Val)
    }
    res := &ListNode{}
    sum := 0
    for len(st1) > 0 || len(st2) > 0 {
        if len(st1) > 0 {
            sum += st1[len(st1)-1]
            st1 = st1[:len(st1)-1]
        }
        if len(st2) > 0 {
            sum += st2[len(st2)-1]
            st2 = st2[:len(st2)-1]
        }
        res.Val = sum % 10
        next := &ListNode{sum/10, nil}
        next.Next = res
        res = next
        sum /= 10
    }
    if res.Val == 0 {
        return res.Next
    }
    return res
}
