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
    int longestConsecutive(TreeNode* root) {
        int ans = 0;
        helper(root, NULL, 0, ans);
        return ans;
    }
    
    void helper(TreeNode* node, TreeNode* parent, int length, int &ans) {
        if (node == NULL)
            return;
        length = (parent && node->val == parent->val + 1) ? length+1 : 1;
        ans = max(ans, length);
        helper(node->left, node, length, ans);
        helper(node->right, node, length, ans);
    }
};
