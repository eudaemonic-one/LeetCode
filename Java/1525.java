class Solution {
    public int numSplits(String s) {
        if (s.length() <= 1) {
            return 0;
        }
        int res = 0;
        Map<Character, Integer> left = new HashMap<>();
        left.put(s.charAt(0), 1);
        Map<Character, Integer> right = new HashMap<>();
        for (int i = 1; i < s.length(); i++) {
            char ch = s.charAt(i);
            right.put(ch, right.getOrDefault(ch, 0) + 1);
        }
        for (int i = 1; i < s.length(); i++) {
            if (left.size() == right.size()) {
                res += 1;
            }
            char ch = s.charAt(i);
            if (right.get(ch) == 1) {
                right.remove(ch);
            } else {
                right.put(ch, right.getOrDefault(ch, 0) - 1);
            }
            left.put(ch, left.getOrDefault(ch, 0) + 1);
        }
        return res;
    }
}