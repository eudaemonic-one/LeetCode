class Solution {
    public int monotoneIncreasingDigits(int N) {
        if (N == 0) {
            return 0;
        }
        char[] s = String.valueOf(N).toCharArray();
        int i = 1;
        // iterate until the first cliff
        while (i < s.length && s[i-1] <= s[i]) {
            i++;
        }
        // truncate after cliff
        while (i > 0 && i < s.length && s[i-1] > s[i]) {
            s[i-1] -= 1;
            i -= 1;
        }
        while (++i < s.length) {
            s[i] = '9';
        }
        return Integer.parseInt(String.valueOf(s));
    }
}