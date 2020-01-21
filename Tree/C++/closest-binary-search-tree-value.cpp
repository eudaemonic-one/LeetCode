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
    int closestValue(TreeNode* root, double target) {
        int closest = root->val, val = root->val;
        while (root) {
            val = root->val;
            closest = (abs(val-target) < abs(closest-target)) ? val : closest;
            root = (target < root->val) ? root->left : root->right;
        }
        return closest;
    }
};
