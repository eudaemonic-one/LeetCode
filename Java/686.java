class Solution {
    public int repeatedStringMatch(String a, String b) {
        StringBuilder sb = new StringBuilder(a);
        int repeat = 1;
        for (; sb.length() < b.length(); repeat++) {
            sb.append(a);
        }
        if (sb.indexOf(b) >= 0) {
            return repeat;
        }
        if (sb.append(a).indexOf(b) >= 0) {
            return repeat + 1;
        }
        return -1;
    }
}