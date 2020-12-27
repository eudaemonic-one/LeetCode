class Solution {
    public List<List<Integer>> reconstructMatrix(int upper, int lower, int[] colsum) {
        int n = colsum.length;
        int[][] mat = new int[2][n];
        for (int j = 0; j < n; j++) {
            if (colsum[j] == 2 || (colsum[j] == 1 && upper >= lower)) {
                mat[0][j] = 1;
            }
            if (colsum[j] == 2 || (colsum[j] == 1 && mat[0][j] == 0)) {
                mat[1][j] = 1;
            }
            upper -= mat[0][j];
            lower -= mat[1][j];
        }
        List<List<Integer>> res = new ArrayList<>();
        for (int i = 0; i < 2; i++) {
            List<Integer> row = new ArrayList<>();
            for (int j = 0; j < n; j++) {
                row.add(mat[i][j]);
            }
            res.add(row);
        }
        if (upper == 0 && lower == 0) {
            return res;
        }
        return new ArrayList<>();
    }
}