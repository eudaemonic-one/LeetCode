/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/
class Solution {
public:
    Node* connect(Node* root) {
        if (root == NULL) {
            return NULL;
        }
        Node *leftmost = root, *prev, *curr;
        while (leftmost) {
            prev = NULL;
            curr = leftmost;
            leftmost = NULL;
            while (curr) {
                processChild(curr->left, prev, leftmost);
                processChild(curr->right, prev, leftmost);
                curr = curr->next;
            }
        }
        return root;
    }
    
    void processChild(Node *child, Node *&prev, Node *&leftmost) {
        if (child) {
            if (prev) {
                prev->next = child;
            } else {
                leftmost = child;
            }
            prev = child;
        }
    }
};
