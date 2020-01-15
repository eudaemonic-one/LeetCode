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
        queue<TreeNode*> q;
        q.push(root);
        TreeNode * node;
        int depth = 0;
        while (!q.empty()) {
            int n = q.size();
            ++depth;
            for (int i = 0; i < n; i++) {
                node = q.front();
                q.pop();
                if (node->left == NULL && node->right == NULL) {
                    return depth;
                }
                if (node->left) {
                    q.push(node->left);
                }
                if (node->right) {
                    q.push(node->right);
                }
            }
        }
        return depth;
    }
};
