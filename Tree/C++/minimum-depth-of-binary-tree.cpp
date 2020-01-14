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
        int min_depth = INT_MAX;
        stack<pair<TreeNode*, int>> stack;
        stack.push(pair<TreeNode*, int>(root, 1));
        while (!stack.empty()) {
            pair<TreeNode*, int> curr = stack.top();
            stack.pop();
            TreeNode* node = curr.first;
            int depth = curr.second;
            if (node->left == NULL && node->right == NULL) {
                min_depth = min(min_depth, depth);
            }
            if (node->left) {
                stack.push(pair<TreeNode*, int>(node->left, depth+1));
            }
            if (node->right) {
                stack.push(pair<TreeNode*, int>(node->right, depth+1));
            }
        }
        return min_depth;
    }
};
