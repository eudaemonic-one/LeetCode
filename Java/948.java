// Approach #2: Greedy
class Solution {
    public int bagOfTokensScore(int[] tokens, int P) {
        Arrays.sort(tokens);
        int lo = 0, hi = tokens.length - 1;
        int score = 0, ans = 0;
        while (lo <= hi && (P >= tokens[lo] || score >= 1)) {
            while (lo <= hi && P >= tokens[lo]) {
                P -= tokens[lo++];
                ++score;
            }
            ans = Math.max(ans, score);
            if (lo <= hi && score > 0) {
                P += tokens[hi--];
                --score;
            }
        }
        return ans;
    }
}

// Approach #1: Backtracking (Time Limited Exceeded)
// class Solution {
//     public int bagOfTokensScore(int[] tokens, int P) {
//         int n = tokens.length;
//         boolean[] used = new boolean[n];
//         int[] res = new int[1];
//         dfs(tokens, used, P, 0, res);
//         return res[0];
//     }
    
//     private void dfs(int[] tokens, boolean[] used, int power, int score, int[] res) {
//         res[0] = Math.max(res[0], score);
//         for (int i = 0; i < tokens.length; i++) {
//             if (!used[i]) {
//                 used[i] = true;
//                 if (power >= tokens[i]) {
//                     dfs(tokens, used, power - tokens[i], score + 1, res);
//                 }
//                 if (score >= 1) {
//                     dfs(tokens, used, power + tokens[i], score - 1, res);
//                 }
//                 used[i] = false;
//             }
//         }
//     }
// }