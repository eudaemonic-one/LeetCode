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
    vector<int> rightSideView(TreeNode* root) {
        vector<int> ans;
        queue<TreeNode*> q;
        TreeNode *node;
        int n = 0;
        if (root == NULL) {
            return ans;
        }
        q.push(root);
        while (!q.empty()) {
            n = q.size();
            for (int i = 0; i < n; i++) {
                node = q.front();
                q.pop();
                if (node->left) {
                    q.push(node->left);
                }
                if (node->right) {
                    q.push(node->right);
                }
            }
            ans.push_back(node->val);
        }
        return ans;
    }
};
