class Solution {
    public int maxDistToClosest(int[] seats) {
        int n = seats.length;
        int prev = -1, next = 0;
        int res = 0;
        for (int i = 0; i < n; i++) {
            if (seats[i] == 1) {
                prev = i;
            } else {
                while (next < n && seats[next] == 0 || next < i) {
                    next++;
                }
                int left = prev == -1 ? n : i - prev;
                int right = next == n ? n : next - i;
                res = Math.max(res, Math.min(left, right));
            }
        }
        return res;
    }
}