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
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> res;
        TreeNode *node;
        stack<TreeNode*> s;
        if (root == NULL) {
            return res;
        }
        s.push(root);
        while (!s.empty()) {
            node = s.top();
            s.pop();
            res.push_back(node->val);
            if (node->left) {
                s.push(node->left);
            }
            if (node->right) {
                s.push(node->right);
            }
        }
        reverse(res.begin(), res.end());
        return res;
    }
};
