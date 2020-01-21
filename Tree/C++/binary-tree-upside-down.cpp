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
    TreeNode* upsideDownBinaryTree(TreeNode *root) {
        map<TreeNode*, TreeNode*> parents;
        TreeNode *newRoot, *parent;
        while (root) {
            if (root->left) {
                parents[root->left] = root;
                root = root->left;
            } else {
                break;
            }
        }
        newRoot = root;
        while (root) {
            if (parents[root]) {
                parent = parents[root];
                root->left = parent->right;
                root->right = parent;
                root = parent;
            } else {
                root->left = NULL;
                root->right = NULL;
                break;
            }
        }
        return newRoot;
    }
};
