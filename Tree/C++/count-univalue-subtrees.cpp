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
        helper(root, 0, ans);
        return ans;
    }
    
    bool helper(TreeNode *root, int parent, int &ans) {
        bool left, right;
        if (root == NULL)
            return true;
        left = helper(root->left, root->val, ans);
        right = helper(root->right, root->val, ans);
        if (!(left && right))
            return false;
        ++ans;
        return root->val == parent;
    }
};
