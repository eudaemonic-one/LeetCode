class Solution {
    private static final int[] POW_10 = new int[]{1, 10, 100, 1000, 10000, 100000, 1000000};
    
    public boolean isSolvable(String[] words, String result) {
        int[] charCount = new int[91];
        boolean[] nonLeadingZero = new boolean[91];
        Set<Character> charSet = new HashSet<>();
        for (String word : words) {
            char[] cs = word.toCharArray();
            for (int i = 0; i < cs.length; i++) {
                if (i == 0 && cs.length > 1) {
                    nonLeadingZero[cs[i]] = true;
                }
                charSet.add(cs[i]);
                charCount[cs[i]] += POW_10[cs.length - i - 1];
            }
        }
        char[] cs = result.toCharArray();
        for (int i = 0; i < cs.length; i++) {
            if (i == 0 && cs.length > 1) {
                nonLeadingZero[cs[i]] = true;
            }
            charSet.add(cs[i]);
            charCount[cs[i]] -= POW_10[cs.length - i - 1];
        }
        boolean[] used = new boolean[10];
        char[] charList = new char[charSet.size()];
        int i = 0;
        for (char c : charSet) {
            charList[i++] = c;
        }
        return backtracking(used, charList, nonLeadingZero, 0, 0, charCount);
    }
    
    private boolean backtracking(boolean[] used, char[] charList, boolean[] nonLeadingZero, int step, int diff, int[] charCount) {
        if (step == charList.length) {
            return diff == 0;
        }
        for (int d = 0; d <= 9; d++) {
            char c = charList[step];
            if (!used[d] && (d > 0 || !nonLeadingZero[c])) {
                used[d] = true;
                if (backtracking(used, charList, nonLeadingZero, step + 1, diff + charCount[c] * d, charCount)) {
                    return true;
                }
                used[d] = false;
            }
        }
        return false;
    }
}