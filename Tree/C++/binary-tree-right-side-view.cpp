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
        map<int, int> depthMap;
        stack<pair<int, TreeNode*>> s;
        TreeNode *node;
        int depth = 0, maxDepth = 0;
        if (root == NULL) {
            return ans;
        }
        s.push(pair<int, TreeNode*>(0, root));
        while (!s.empty()) {
            depth = s.top().first;
            node = s.top().second;
            s.pop();
            if (node) {
                maxDepth = max(maxDepth, depth);
                if (depthMap.find(depth) == depthMap.end()) {
                    depthMap[depth] = node->val;
                }
                s.push(pair<int,TreeNode*>(depth+1, node->left));
                s.push(pair<int,TreeNode*>(depth+1, node->right));
            }
        }
        for (int i = 0; i <= maxDepth; i++) {
            ans.push_back(depthMap[i]);
        }
        return ans;
    }
};
