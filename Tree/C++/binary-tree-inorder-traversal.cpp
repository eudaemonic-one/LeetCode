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
        helper(root, &res);
        return res;
    }
    
    void helper(TreeNode* node, vector<int>* res) {
        if (node == NULL) {
            return;
        }
        helper(node->left, res);
        (*res).push_back(node->val);
        helper(node->right, res);
    }
};
