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
    int closestValue(TreeNode* root, double target) {
        long prev = LONG_MIN;
        stack<TreeNode*> s;
        while ((!s.empty()) || (root != NULL)) {
            while (root) {
                s.push(root);
                root = root->left;
            }
            root = s.top();
            s.pop();
            if (prev <= target && target < root->val)
                return abs(prev-target) < abs(root->val-target) ? prev : root->val;
            prev = root->val;
            root = root->right;
        }
        return int(prev);
    }
};
