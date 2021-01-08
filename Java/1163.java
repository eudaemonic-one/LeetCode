class Solution {
    public String lastSubstring(String s) {
        int i = 0, j = 1, k = 0, n = s.length();
        while (j + k < n) {
            if (s.charAt(i+k) == s.charAt(j+k)) {
                k++;
            } else if (s.charAt(i+k) > s.charAt(j+k)) {
                j = j + k + 1;
                k = 0;
            } else {
                i = Math.max(j, i + k + 1);
                j = i + 1;
                k = 0;
            }
        }
        return s.substring(i);
    }
}