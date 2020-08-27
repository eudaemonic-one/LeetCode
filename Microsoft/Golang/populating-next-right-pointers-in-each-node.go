/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            node := queue[i]
            if i < n-1 {
                node.Next = queue[i+1]
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[n:]
    }
    return root
}
