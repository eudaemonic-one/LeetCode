class Solution {
    public int lengthLongestPath(String input) {
        int maxLen = 0, curLen = 0;
        String[] tokens = input.split("\n");
        Stack<Integer> stack = new Stack<>();
        for (String token : tokens) {
            int level = 0;
            boolean isFile = false;
            for (int i = 0; i < token.length(); i++) {
                if (token.charAt(i) == '\t') {
                    ++level;
                } else if (token.charAt(i) == '.') {
                    isFile = true;
                }
            }
            
            while (stack.size() > level) {
                curLen -= stack.pop();
            }
            
            int tokenLen = token.length() - level + 1;
            curLen += tokenLen;
            if (isFile) {
                if (curLen - 1 > maxLen) {
                    maxLen = curLen - 1;
                }
            }
            stack.add(tokenLen);
        }
        return maxLen;
    }
}