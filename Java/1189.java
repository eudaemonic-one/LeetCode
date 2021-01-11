class Solution {
    public int maxNumberOfBalloons(String text) {
        Map<Character, Integer> dict = new HashMap<>();
        for (char ch: text.toCharArray()) {
            dict.put(ch, dict.getOrDefault(ch, 0) + 1);
        }
        int res = Integer.MAX_VALUE;
        res = Math.min(res, dict.getOrDefault('b', 0));
        res = Math.min(res, dict.getOrDefault('a', 0));
        res = Math.min(res, dict.getOrDefault('l', 0) / 2);
        res = Math.min(res, dict.getOrDefault('o', 0) / 2);
        res = Math.min(res, dict.getOrDefault('n', 0));
        return res;
    }
}