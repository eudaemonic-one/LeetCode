class Solution {    
    public int numDistinctIslands(int[][] grid) {
        boolean[][] visited = new boolean[grid.length][grid[0].length];
        Set<List<Integer>> shapes = new HashSet<>();
        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[0].length; c++) {
                List<Integer> shape = new ArrayList<>();
                helper(grid, visited, shape, r, c, 0);
                if (!shape.isEmpty()) {
                    shapes.add(shape);
                }
            }
        }
        return shapes.size();
    }
    
    private void helper(int[][] grid, boolean[][] visited, List<Integer> shape, int r, int c, int dir) {
        if (0 <= r && r < grid.length && 0 <= c && c < grid[0].length && !visited[r][c] && grid[r][c] == 1) {
            visited[r][c] = true;
            shape.add(dir);
            helper(grid, visited, shape, r+1, c, 1);
            helper(grid, visited, shape, r-1, c, 2);
            helper(grid, visited, shape, r, c+1, 3);
            helper(grid, visited, shape, r, c-1, 4);
            shape.add(0);
        }
    }
}