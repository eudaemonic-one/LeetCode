class Solution {
    public int[] kWeakestRows(int[][] mat, int k) {
        int m = mat.length;
        int n = mat[0].length;
        int[][] pairs = new int[m][2];
        for (int i = 0; i < m; i++) {
            int count = 0;
            for (int j = 0; j < n; j++) {
                if (mat[i][j] == 1) {
                    count += 1;
                }
            }
            pairs[i][0] = i;
            pairs[i][1] = count;
        }
        Arrays.sort(pairs, (a, b) -> {
            if (a[1] == b[1]) {
                return a[0] - b[0];
            }
            return a[1] - b[1];
        });
        int[] res = new int[k];
        for (int i = 0; i < k; i++) {
            res[i] = pairs[i][0];
        }
        return res;
    }
}