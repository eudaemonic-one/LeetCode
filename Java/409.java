class Solution {
    public int longestPalindrome(String s) {
        int[] counter = new int[128];
        for (char ch : s.toCharArray()) {
            counter[ch]++;
        }
        int ans = 0;
        for (int count : counter) {
            ans += count / 2 * 2;
            if (ans % 2 == 0 && count % 2 == 1) {
                ans++;
            }
        }
        return ans;
    }
}