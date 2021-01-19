class Solution {
    private Set<Integer> seen;
    private List<Set<Integer>> regions;
    private List<Set<Integer>> frontiers;
    private List<Integer> perimeters;
    private static int[] dr = new int[]{0, 1, 0 , -1};
    private static int[] dc = new int[]{-1, 0, 1, 0};
    
    public int containVirus(int[][] grid) {
        int R = grid.length;
        int C = grid[0].length;
        int res = 0;
        while (true) {
            seen = new HashSet<>();
            regions = new ArrayList<>();
            frontiers = new ArrayList<>();
            perimeters = new ArrayList<>();
            
            for (int r = 0; r < R; r++) {
                for (int c = 0; c < C; c++) {
                    if (grid[r][c] == 1 && !seen.contains(r*C+c)) {
                        regions.add(new HashSet<>());
                        frontiers.add(new HashSet<>());
                        perimeters.add(0);
                        dfs(grid, R, C, r, c);
                    }
                }
            }
            
            if (regions.isEmpty()) {
                break;
            }
            int triageIndex = 0;
            for (int i = 1; i < frontiers.size(); i++) {
                if (frontiers.get(triageIndex).size() < frontiers.get(i).size()) {
                    triageIndex = i;
                }
            }
            res += perimeters.get(triageIndex);
            
            for (int i = 0; i < regions.size(); i++) {
                if (i == triageIndex) {
                    for (int code: regions.get(i)) {
                        grid[code/C][code%C] = -1;
                    }
                } else {
                    for (int code: regions.get(i)) {
                        int r = code / C, c = code % C;
                        for (int k = 0; k < 4; k++) {
                            int nr = r + dr[k], nc = c + dc[k];
                            if (0 <= nr && nr < R && 0 <= nc && nc < C && grid[nr][nc] == 0) {
                                grid[nr][nc] = 1;
                            }
                        }
                    }
                }
            }
        }
        return res;
    }
    
    private void dfs(int[][] grid, int R, int C, int r, int c) {
        int coor = r*C + c;
        if (!seen.contains(coor)) {
            seen.add(coor);
            int n = regions.size();
            regions.get(n-1).add(coor);
            for (int k = 0; k < 4; k++) {
                int nr = r + dr[k], nc = c + dc[k];
                if (0 <= nr && nr < R && 0 <= nc && nc < C) {
                    if (grid[nr][nc] == 1) {
                        dfs(grid, R, C, nr, nc);
                    } else if (grid[nr][nc] == 0) {
                        frontiers.get(n-1).add(nr*C+nc);
                        perimeters.set(n-1, perimeters.get(n-1) + 1);
                    }
                }
            }
        }
    }
}