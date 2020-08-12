/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    table := make(map[*Node]*Node)
    for curr := head; curr != nil; curr = curr.Next {
        table[curr] = &Node{curr.Val, nil, nil}
    }
    for curr := head; curr != nil; curr = curr.Next {
        table[curr].Next = table[curr.Next]
        table[curr].Random = table[curr.Random]
    }
    return table[head]
}
