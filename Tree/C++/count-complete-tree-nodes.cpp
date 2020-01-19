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
    int countNodes(TreeNode* root) {
        int depth = 0, left, pivot, right;
        TreeNode *node = root;
        if (root == NULL)
            return 0;
        while (node && node->left) {
            node = node->left;
            ++depth;
        }
        if (depth == 0)
            return 1;
        left = 1, right = pow(2, depth)-1;
        while (left <= right) {
            pivot = floor((left + right) / 2);
            if (exists(pivot, depth, root))
                left = pivot + 1;
            else
                right = pivot - 1;
        }
        return int(pow(2, depth)) - 1 + left;
    }
    
    bool exists(int idx, int depth, TreeNode *root) {
        int left = 0, right = int(pow(2, depth)) - 1, pivot;
        for (int i = 0; i < depth; i++) {
            pivot = floor((left + right) / 2);
            if (idx <= pivot) {
                root = root->left;
                right = pivot;
            } else {
                root = root->right;
                left = pivot + 1;
            }
        }
        return root != NULL;
    }
};
