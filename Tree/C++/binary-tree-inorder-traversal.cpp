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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        TreeNode *node = root;
        if (!root) {
            return res;
        }
        stack<TreeNode *> stack;
        while (!stack.empty() || node) {
            while (node) {
                stack.push(node);
                node = node->left;
            }
            if (stack.empty()) {
                break;
            }
            node = stack.top();
            stack.pop();
            res.push_back(node->val);
            node = node->right;
        }
        return res;
    }
};
