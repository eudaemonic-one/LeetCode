class Solution {
    public int numSplits(String s) {
        int n = s.length();
        if (n <= 1) {
            return 0;
        }
        int[] pre = new int[n];
        int[] suf = new int[n];
        Map<Character, Integer> freq = new HashMap<>();
        for (int i = 0; i < n; i++) {
            char ch = s.charAt(i);
            freq.put(ch, freq.getOrDefault(ch, 0) + 1);
            pre[i] = freq.size();
        }
        freq.clear();
        for (int i = n-1; i >= 0; i--) {
            char ch = s.charAt(i);
            freq.put(ch, freq.getOrDefault(ch, 0) + 1);
            suf[i] = freq.size();
        }
        int res = 0;
        for (int i = 1; i < n; i++) {
            if (pre[i-1] == suf[i]) {
                res += 1;
            }
        }
        return res;
    }
}

// Approach 1: Double Scan
// class Solution {
//     public int numSplits(String s) {
//         if (s.length() <= 1) {
//             return 0;
//         }
//         int res = 0;
//         Map<Character, Integer> left = new HashMap<>();
//         left.put(s.charAt(0), 1);
//         Map<Character, Integer> right = new HashMap<>();
//         for (int i = 1; i < s.length(); i++) {
//             char ch = s.charAt(i);
//             right.put(ch, right.getOrDefault(ch, 0) + 1);
//         }
//         for (int i = 1; i < s.length(); i++) {
//             if (left.size() == right.size()) {
//                 res += 1;
//             }
//             char ch = s.charAt(i);
//             if (right.get(ch) == 1) {
//                 right.remove(ch);
//             } else {
//                 right.put(ch, right.getOrDefault(ch, 0) - 1);
//             }
//             left.put(ch, left.getOrDefault(ch, 0) + 1);
//         }
//         return res;
//     }
// }