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
    TreeNode* removeLeafNodes(TreeNode* root, int target) {
        if (root == NULL) {
            return NULL;
        }
        helper(root, target);
        if ((root->val == target) && (root->left == NULL) && (root->right == NULL)) {
            return NULL;
        }
        return root;
    }
    
    bool helper(TreeNode* root, int target) {
        bool flag = false;
        if (root == NULL) {
            return false;
        }
        if (root->left) {
            if (helper(root->left, target)) {
                root->left = NULL;
            }
        }
        if (root->right) {
            if (helper(root->right, target)) {
                root->right = NULL;
            }
        }
        if ((root->val == target) && (root->left == NULL) && (root->right == NULL)) {
            root = NULL;
            flag = true;
        }
        return flag;
    }
};