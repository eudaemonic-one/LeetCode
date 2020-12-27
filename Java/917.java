class Solution {
    public String reverseOnlyLetters(String S) {
        int l = 0;
        int r = S.length() - 1;
        if (l >= r) {
            return S;
        }
        char[] bytes = S.toCharArray();
        while (l < r) {
            while (l < r && !Character.isLetter(bytes[l])) {
                ++l;
            }
            while (l < r && !Character.isLetter(bytes[r])) {
                --r;
            }
            char ch = bytes[l];
            bytes[l] = bytes[r];
            bytes[r] = ch;
            ++l;
            --r;
        }
        return String.valueOf(bytes);
    }
}