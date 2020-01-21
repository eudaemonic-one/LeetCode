/**
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
    TreeNode* inorderSuccessor(TreeNode* root, TreeNode* p) {
        if (p->right) {
            p = p->right;
            while (p->left) {
                p = p->left;
            }
            return p;
        }
        int prev = INT_MIN;
        stack<TreeNode*> s;
        while (!s.empty() || root) {
            while (root) {
                s.push(root);
                root = root->left;
            }
            root = s.top();
            s.pop();
            if (prev == p->val)
                return root;
            prev = root->val;
            root = root->right;
        }
        return NULL;
    }
};
