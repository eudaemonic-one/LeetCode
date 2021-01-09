class Solution {
    public int countBinarySubstrings(String s) {
        int zero = 0, one = 0, res = 0;
        char lastCh = ' ';
        for (char ch: s.toCharArray()) {
            if (zero > 0 && one > 0 && ch != lastCh) {
                res += Math.min(zero, one);
                if (lastCh == '0') {
                    one = 0;
                } else {
                    zero = 0;
                }
            }
            if (ch == '0') {
                zero++;
            } else {
                one++;
            }
            lastCh = ch;
        }
        res += Math.min(zero, one);
        return res;
    }
}