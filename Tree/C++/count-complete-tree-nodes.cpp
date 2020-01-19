o/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int countNodes(TreeNode* root) {
        int count = 0;
        TreeNode *node;
        stack<TreeNode*> s;
        s.push(root);
        while (!s.empty()) {
            node = s.top();
            s.pop();
            if (node) {
                ++count;
                if (node->left) {
                    s.push(node->left);
                }
                if (node->right) {
                    s.push(node->right);
                }
            }
        }
        return count;
    }
};
