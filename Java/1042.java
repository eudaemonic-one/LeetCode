class Solution {
    public int[] gardenNoAdj(int n, int[][] paths) {
        Map<Integer, Set<Integer>> graph = new HashMap<>();
        for (int i = 1; i <= n; i++) {
            graph.put(i, new HashSet<>());
        }
        for (int[] path : paths) {
            graph.get(path[0]).add(path[1]);
            graph.get(path[1]).add(path[0]);
        }
        int[] res = new int[n];
        for (int i = 1; i <= n; i++) {
            int[] colors = new int[5];
            for (int v : graph.get(i)) {
                colors[res[v-1]] = 1;
            }
            for (int c = 4; c >= 0; c--) {
                if (colors[c] == 0) {
                    res[i-1] = c;
                    break;
                }
            }
        }
        return res;
    }
}