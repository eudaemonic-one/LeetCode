class Solution {
    public int getMaximumGenerated(int n) {
        if (n <= 1) {
            return n;
        }
        int res = 1;
        int[] nums = new int[n+1];
        nums[1] = 1;
        for (int i = 1; i <= n/2; i++) {
            nums[i*2] = nums[i];
            if (i*2 + 1 <= n) {
                nums[2*i+1] = nums[i] + nums[i+1];
                res = Math.max(res, nums[2*i+1]);
            }
        }
        return res;
    }
}