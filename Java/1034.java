class Solution {
    private static int[] dr = new int[]{0, -1, 0, 1};
    private static int[] dc = new int[]{-1, 0, 1, 0};
    
    public int[][] colorBorder(int[][] grid, int r0, int c0, int color) {
        dfs(grid, r0, c0, grid[r0][c0]);
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[0].length; j++) {
                if (grid[i][j] < 0) {
                    grid[i][j] = color;
                }
            }
        }
        return grid;
    }
    
    private void dfs(int[][] grid, int r, int c, int color) {
        grid[r][c] = -color;
        int count = 0;
        for (int i = 0; i < 4; i++) {
            int nr = r + dr[i];
            int nc = c + dc[i];
            if (nr < 0 || nr >= grid.length || nc < 0 || nc >= grid[0].length || Math.abs(grid[nr][nc]) != color) {
                continue;
            }
            count++;
            if (grid[nr][nc] == color) {
                dfs(grid, nr, nc, color);
            }
        }
        if (count == 4) {
            grid[r][c] = color;
        }
    }
}