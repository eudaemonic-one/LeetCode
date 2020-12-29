class Solution {
    public String makeGood(String s) {
        if (s.length() <= 1) {
            return s;
        }
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char ch = s.charAt(i);
            if (!stack.isEmpty() && Math.abs(stack.peek() - ch) == 32) {
                stack.pop();
            } else {
                stack.push(ch);
            }
        }
        char[] res = new char[stack.size()];
        int idx = stack.size() - 1;
        while (!stack.isEmpty()) {
            res[idx--] = stack.pop();
        }
        return String.valueOf(res);
    }
}