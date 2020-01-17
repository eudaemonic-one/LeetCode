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
        helper(root, res);
        return res;
    }
    
    void helper(TreeNode* root, vector<int> &output) {
        if (root == NULL) {
            return;
        }
        helper(root->left, output);
        helper(root->right, output);
        output.push_back(root->val);
    }
};
