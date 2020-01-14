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
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> res;
        int level = 0;
        queue<TreeNode*> q;
        q.push(root);
        if (root == NULL) {
            return res;
        }
        while (!q.empty()) {
            res.push_back(vector<int>());
            int n = q.size();
            for (int i = 0; i < n; i++) {
                TreeNode *node = q.front();
                if (node->left) {
                    q.push(node->left);
                }
                if (node->right) {
                    q.push(node->right);
                }
                res[level].push_back(node->val);
                q.pop();
            }
            if (level % 2 == 1) {
                reverse(res[level].begin(), res[level].end());
            }
            ++level;
        }
        return res;
    }
};
