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
    int maxPathSum(TreeNode* root) {
        int res = INT_MIN;
        if (root == NULL) {
            return 0;
        }
        helper(root, res);
        return res;
    }
    
    int helper(TreeNode *root, int &res) {
        int left, right, tmp;
        if (root == NULL) {
            return 0;
        }
        left = helper(root->left, res);
        right = helper(root->right, res);
        tmp = max(
            max(root->val, left+root->val+right),
            max(left+root->val, right+root->val)
        );
        res = max(res, tmp);
        return max(root->val, max(left+root->val, right+root->val));
    }
};
