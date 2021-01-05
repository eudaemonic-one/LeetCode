class Solution {
    public String findLexSmallestString(String s, int a, int b) {
        Set<String> visited = new HashSet<>();
        String[] res = new String[]{s};
        dfs(res, s, a, b, visited);
        return res[0];
    }
    
    private void dfs(String[] res, String s, int a, int b, Set<String> visited) {
        if (visited.contains(s)) {
            return;
        }
        visited.add(s);
        if (s.compareTo(res[0]) < 0) {
            res[0] = s;
        }
        dfs(res, add(s, a), a, b, visited);
        dfs(res, rotate(s, b), a, b, visited);
    }
    
    private String add(String s, int a) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < s.length(); i++) {
            if (i % 2 == 0) {
                sb.append(s.charAt(i));
            } else {
                sb.append((s.charAt(i) - '0' + a) % 10);
            }
        }
        return sb.toString();
    }
    
    private String rotate(String s, int b) {
        int n = s.length();
        return s.substring(n-b, n) + s.substring(0, n-b);
    }
}