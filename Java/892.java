class Solution {
    public int surfaceArea(int[][] grid) {
        int res = 0;
        int n = grid.length;
        int[] dr = new int[]{0, 1, 0, -1};
        int[] dc = new int[]{1, 0, -1, 0};
        for (int r = 0; r < n; r++) {
            for (int c = 0; c < n; c++) {
                if (grid[r][c] > 0) {
                    res += 2;
                    for (int i = 0; i < 4; i++) {
                        int nr = r + dr[i];
                        int nc = c + dc[i];
                        int nv = 0;
                        if (0 <= nr && nr < n && 0 <= nc && nc < n) {
                            nv = grid[nr][nc];
                        }
                        res += Math.max(grid[r][c]-nv, 0);
                    }
                }
            }
        }
        return res;
    }
}