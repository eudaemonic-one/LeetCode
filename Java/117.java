/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}
    
    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
    private Node prev;
    private Node leftmost;
    
    public Node connect(Node root) {
        if (root == null) {
            return null;
        }
        leftmost = root;
        Node curr = root;
        while (leftmost != null) {
            prev = null;
            curr = leftmost;
            leftmost = null;
            while (curr != null) {
                processChild(curr.left);
                processChild(curr.right);
                curr = curr.next;
            }
        }
        return root;
    }
    
    private void processChild(Node childNode) {
        if (childNode != null) {
            if (prev != null) {
                prev.next = childNode;
            } else {
                leftmost = childNode;
            }
            prev = childNode;
        }
    }
}