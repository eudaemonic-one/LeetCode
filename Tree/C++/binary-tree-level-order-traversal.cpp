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
    vector<vector<int>> res;
    
    vector<vector<int>> levelOrder(TreeNode* root) {
        helper(root, 0);
        return res;
    }
    
    void helper(TreeNode* root, int level) {
        if (root == NULL) {
            return;
        }
        if (level == res.size()) {
            res.push_back(vector<int>());
        }
        res[level].push_back(root->val);
        helper(root->left, level+1);
        helper(root->right, level+1);
    }
};
