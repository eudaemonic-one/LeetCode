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
    int minDepth(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }
        if (root->left == NULL && root->right == NULL) {
            return 1;
        }
        int min_length = INT_MAX;
        if (root->left) {
            min_length = min(min_length, minDepth(root->left));
        }
        if (root->right) {
            min_length = min(min_length, minDepth(root->right));
        }
        return min_length+1;
    }
};
