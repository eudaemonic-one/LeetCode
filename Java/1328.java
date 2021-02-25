class Solution {
    public String breakPalindrome(String palindrome) {
        if (palindrome.length() <= 1) {
            return "";
        }
        char[] str = palindrome.toCharArray();
        for (int i = 0; i < palindrome.length() / 2; i++) {
            if (str[i] != 'a') {
                str[i] = 'a';
                return String.valueOf(str);
            }
        }
        str[str.length-1] = 'b';
        return String.valueOf(str);
    }
}