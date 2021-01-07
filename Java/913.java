class Solution {
    private final int DRAW = 0;
    private final int MOUSE = 1;
    private final int CAT = 2;
    
    public int catMouseGame(int[][] graph) {
        int n = graph.length;
        int[][][] color = new int[50][50][3];
        int[][][] degree = new int[50][50][3];
        for (int m = 0; m < n; m++) {
            for (int c = 0; c < n; c++) {
                degree[m][c][MOUSE] = graph[m].length;
                degree[m][c][CAT] = graph[c].length;
                for (int node: graph[c]) {
                    if (node == 0) {
                        degree[m][c][CAT]--;
                        break;
                    }
                }
            }
        }
        Queue<int[]> queue = new LinkedList<>();
        for (int i = 0; i < n; i++) {
            for (int t = 1; t <= 2; t++) {
                color[0][i][t] = MOUSE;
                queue.add(new int[]{0, i, t, MOUSE});
                if (i != 0) {
                    color[i][i][t] = CAT;
                    queue.add(new int[]{i, i, t, CAT});
                }
            }
        }
        while (!queue.isEmpty()) {
            int[] node = queue.poll();
            int i = node[0], j = node[1], t = node[2], c = node[3];
            for (int[] parent: parents(graph, i, j, t)) {
                int i2 = parent[0], j2 = parent[1], t2 = parent[2];
                if (color[i2][j2][t2] == DRAW) {
                    if (t2 == c) { // if parent can make a winning move
                        color[i2][j2][t2] = c;
                        queue.add(new int[]{i2, j2, t2, c});
                    } else {
                        degree[i2][j2][t2]--;
                        if (degree[i2][j2][t2] == 0) { // if all children of this parent may lose
                            color[i2][j2][t2] = 3 - t2;
                            queue.add(new int[]{i2, j2, t2, 3-t2});
                        }
                    }
                }
            }
        }
        return color[1][2][1];
    }
    
    private List<int[]> parents(int[][] graph, int m, int c, int t) {
        List<int[]> parents = new ArrayList<>();
        if (t == CAT) { // parent must be MOUSE
            for (int m2: graph[m]) {
                parents.add(new int[]{m2, c, 3-t});
            }
        } else { // parent must be CAT
            for (int c2: graph[c]) {
                if (c2 != 0) { // not a hole
                    parents.add(new int[]{m, c2, 3-t});
                }
            }
        }
        return parents;
    }
}