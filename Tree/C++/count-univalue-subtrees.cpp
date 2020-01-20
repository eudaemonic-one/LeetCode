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
    int countUnivalSubtrees(TreeNode* root) {
        int ans = 0;
        if (root == NULL)
            return 0;
        helper(root, ans);
        return ans;
    }
    
    bool helper(TreeNode *root, int &ans) {
        bool flag = true;
        if ((root->left == NULL) && (root->right == NULL)) {
            ++ans;
            return true;
        }
        if (root->left)
            flag = helper(root->left, ans) && flag && (root->left->val == root->val);
        if (root->right)
            flag = helper(root->right, ans) && flag && (root->right->val == root->val);
        if (flag)
            ++ans;
        return flag;
    }
};
