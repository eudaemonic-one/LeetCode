class Solution {
    public String removeOuterParentheses(String S) {
        StringBuilder sb = new StringBuilder();
        int opened = 0;
        for (char c : S.toCharArray()) {
            if (c == '(' && opened++ > 0) {
                sb.append(c);
            }
            if (c == ')' && opened-- > 1) {
                sb.append(c);
            }
        }
        return sb.toString();
    }
}