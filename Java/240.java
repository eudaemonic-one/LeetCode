class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int r = matrix.length-1, c = 0;
        while (0 <= r && c < matrix[0].length) {
            if (matrix[r][c] > target) {
                --r;
            } else if (matrix[r][c] < target) {
                ++c;
            } else {
                return true;
            }
        }
        return false;
    }
}