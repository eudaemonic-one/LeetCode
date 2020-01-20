class Solution {
public:
    bool verifyPreorder(vector<int>& preorder) {
        int lower_bound = INT_MIN;
        stack<int> s;
        for (auto val : preorder) {
            if (val < lower_bound)
                return false;
            while (!s.empty() && val > s.top()) {
                lower_bound = s.top();
                s.pop();
            }
            s.push(val);
        }
        return true;
    }
};
