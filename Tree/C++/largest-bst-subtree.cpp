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
    int largestBSTSubtree(TreeNode* root) {
        int ans = 0;
        int mini, maxi;
        isBST(root, mini, maxi, ans);
        return ans;
    }
    
    bool isBST(TreeNode* node, int &mini, int &maxi, int &ans) {
        if (node == NULL)
            return true;
        int leftNumber = 0, rightNumber = 0;
        int lMin = INT_MIN, lMax = INT_MAX;
        int rMin = INT_MIN, rMax = INT_MAX;
        bool l = isBST(node->left, lMin, lMax, leftNumber);
        bool r = isBST(node->right, rMin, rMax, rightNumber);
        if (l && r) {
            if ((node->left == NULL || node->val > lMax)
                    && (node->right == NULL || node->val < rMin)) {
                ans = leftNumber + rightNumber + 1;
                mini = node->left ? lMin : node->val;
                maxi = node->right ? rMax : node->val;
                return true;
            }
        }
        ans = max(leftNumber, rightNumber);
        return false;
    }
};
