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
    vector<int> res;
    
    vector<int> inorderTraversal(TreeNode* root) {
        helper(root);
        return res;
    }
    
    void helper(TreeNode* node) {
        if (node == NULL) {
            return;
        }
        helper(node->left);
        res.push_back(node->val);
        helper(node->right);
    }
};
