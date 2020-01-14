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
    vector<vector<int>> pathSum(TreeNode* root, int sum) {
        vector<vector<int>> res;
        vector<int> path;
        if (root == NULL) {
            return res;
        }
        helper(root, path, sum, res);
        return res;
    }
    
    void helper(TreeNode *root, vector<int>& path, int sum, vector<vector<int>>& res) {
        if (root == NULL) {
            return;
        }
        path.push_back(root->val);
        sum -= root->val;
        if (root->left == NULL && root->right == NULL) {
            if (sum == 0) {
                res.push_back(path);
            }
        } else {
            if (root->left) {
                helper(root->left, path, sum, res);
            }
            if (root->right) {
                helper(root->right, path, sum, res);
            }
        }
        path.pop_back();
        sum += root->val;
    }
};
