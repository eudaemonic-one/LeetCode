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
    TreeNode *prev;
    
    void flatten(TreeNode* root) {
        if (root == NULL) {
            return;
        }
        if (root->right) {
            flatten(root->right);
        }
        if (root->left) {
            flatten(root->left);
        }
        root->right = prev;
        root->left = NULL;
        prev = root;
    }
};
