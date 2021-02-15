class Solution {
    public int countCharacters(String[] words, String chars) {
        int[] dict = new int[26];
        for (char ch : chars.toCharArray()) {
            ++dict[ch-'a'];
        }
        int res = 0;
        for (String word : words) {
            int[] counter = dict.clone();
            boolean good = true;
            for (char ch : word.toCharArray()) {
                if (--counter[ch-'a'] < 0) {
                    good = false;
                    break;
                }
            }
            if (good) {
                res += word.length();
            }
        }
        return res;
    }
}