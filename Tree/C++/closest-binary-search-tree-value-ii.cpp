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
    vector<int> closestKValues(TreeNode* root, double target, int k) {
        vector<int> ans;
        stack<int> s1, s2; // predecessors and successors
        inorder(root, target, false, s1);
        inorder(root, target, true, s2);
        while (k-- > 0) {
            if (s1.empty()) {
                ans.push_back(s2.top());
                s2.pop();
            } else if (s2.empty()) {
                ans.push_back(s1.top());
                s1.pop();
            } else if (abs(s1.top()-target) < abs(s2.top()-target)) {
                ans.push_back(s1.top());
                s1.pop();
            } else {
                ans.push_back(s2.top());
                s2.pop();
            }
        }
        return ans;
    }
    
    void inorder(TreeNode* root, double &target, bool reverse, stack<int>& s) {
        if (root == NULL)
            return;
        TreeNode *left = root->left, *right = root->right;
        inorder(reverse?right:left, target, reverse, s);
        if ((!reverse && root->val > target) || (reverse && root->val <= target))
            return;
        s.push(root->val);
        inorder(reverse?left:right, target, reverse, s);
    }
};
