class Solution {
    private int[] parent;
    
    public String smallestStringWithSwaps(String s, List<List<Integer>> pairs) {
        if (s == null || s.length() == 0) {
            return null;
        }
        int n = s.length();
        parent = new int[n];
        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }
        for (List<Integer> pair : pairs) {
            union(pair.get(0), pair.get(1));
        }
        Map<Integer, PriorityQueue<Character>> map = new HashMap<>();
        for (int i = 0; i < n; i++) {
            int root = find(i);
            map.putIfAbsent(root, new PriorityQueue<>());
            map.get(root).offer(s.charAt(i));
        }
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < n; i++) {
            sb.append(map.get(find(i)).poll());
        }
        return sb.toString();
    }
    
    private int find(int index) {
        while (parent[index] != index) {
            parent[index] = parent[parent[index]];
            index = parent[index];
        }
        return index;
    }
    
    private void union(int x, int y) {
        int xRoot = find(x);
        int yRoot = find(y);
        if (xRoot < yRoot) {
            parent[yRoot] = xRoot;
        } else {
            parent[xRoot] = yRoot;
        }
    }
}