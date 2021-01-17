class Solution {
    public boolean isPerfectSquare(int num) {
        if (num <= 1) {
            return true;
        }
        long l = 2, r = num / 2;
        while (l <= r) {
            long m = l + (r-l) / 2;
            long square = m * m;
            if (square == num) {
                return true;
            }
            if (square > num) {
                r = m - 1;
            } else {
                l = m + 1;
            }
        }
        return false;
    }
}