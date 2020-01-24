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
    int sumOfLeftLeaves(TreeNode* root) {
        int sum = 0;
        helper(root, NULL, sum);
        return sum;
    }
    
    void helper(TreeNode* root, TreeNode* parent, int &sum) {
        if (root == NULL)
            return;
        if (parent && (root == parent->left)
                && (root->left == NULL) && (root->right == NULL))
            sum += root->val;
        helper(root->left, root, sum);
        helper(root->right, root, sum);
    }
};
