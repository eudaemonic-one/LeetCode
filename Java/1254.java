class Solution {
    private static int[] dr = new int[]{0, 1, 0, -1};
    private static int[] dc = new int[]{-1, 0, 1, 0};
    
    public int closedIsland(int[][] grid) {
        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[0].length; c++) {
                if (r == 0 || c == 0 || r == grid.length-1 || c == grid[0].length-1) {
                    fill(grid, r, c);
                }
            }
        }
        int res = 0;
        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[0].length; c++) {
                if (grid[r][c] == 0) {
                    fill(grid, r, c);
                    ++res;
                }
            }
        }
        return res;
    }
    
    private void fill(int[][] grid, int r, int c) {
        if (r < 0 || r >= grid.length || c < 0 || c >= grid[0].length || grid[r][c] == 1) {
            return;
        }
        grid[r][c] = 1;
        fill(grid, r+1, c);
        fill(grid, r-1, c);
        fill(grid, r, c+1);
        fill(grid, r, c-1);
    }
}