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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> res;
        TreeNode *node = root;
        TreeNode* predecessor;
        if (!root) {
            return res;
        }
        while (node) {
            if (node->left) {
                predecessor = node->left;
                while (predecessor->right && predecessor->right != node) {
                    predecessor = predecessor->right;
                }
                if (!predecessor->right) {
                    predecessor->right = node;
                    node = node->left;
                } else {
                    predecessor->right = NULL;
                    res.push_back(node->val);
                    node = node->right;
                }
                predecessor = NULL;
            } else {
                res.push_back(node->val);
                node = node->right;
            }
        }
        return res;
    }
};
