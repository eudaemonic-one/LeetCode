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
    int minDiffInBST(TreeNode* root) {
        int minDiff = INT_MAX;
        int last = INT_MAX;
        stack<TreeNode*> s;
        while (!s.empty() || root) {
            while (root) {
                s.push(root);
                root = root->left;
            }
            root = s.top();
            s.pop();
            if (last != INT_MAX && (root->val - last < minDiff)) {
                minDiff = root->val - last;
            }
            last = root->val;
            root = root->right;
        }
        return minDiff;
    }
};
