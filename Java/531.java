class Solution {
    public int findLonelyPixel(char[][] picture) {
        int res = 0;
        int m = picture.length;
        int n = picture[0].length;
        int[] row = new int[m];
        int[] col = new int[n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (picture[i][j] == 'B') {
                    row[i]++;
                    col[j]++;
                }
            }
        }
        for (int i = 0; i < m; i++) {
            if (row[i] != 1) {
                continue;
            }
            for (int j = 0; j < n; j++) {
                if (picture[i][j] == 'B') {
                    if (col[j] == 1) {
                        res++;
                    }
                    break;
                }
            }
        }
        return res;
    }
}