class Solution {
    public int numberOfPatterns(int m, int n) {
        int res = 0;
        boolean[] used = new boolean[9];
        for (int step = m; step <= n; step++) {
            res += dfs(used, -1, step);
            for (int i = 0; i < 9; i++) {
                used[i] = false;
            }
        }
        return res;
    }
    
    private int dfs(boolean[] used, int prev, int step) {
        if (step == 0) {
            return 1;
        }
        int sum = 0;
        for (int curr = 0; curr < 9; curr++) {
            if (isValid(used, prev, curr)) {
                used[curr] = true;
                sum += dfs(used, curr, step-1);
                used[curr] = false;
            }
        }
        return sum;
    }
    
    private boolean isValid(boolean[] used, int prev, int curr) {
        if (prev == -1) {
            return true;
        }
        if (used[curr]) {
            return false;
        }
        if ((prev + curr) % 2 == 1) { // adjacent cells
            return true;
        }
        int mid = (prev + curr) / 2;
        if (mid == 4) {
            return used[mid];
        }
        if ((curr % 3 != prev % 3) && (curr / 3 != prev / 3)) {
            return true;
        }
        return used[mid];
    }
}