class Solution {
    public int splitArray(int[] nums, int m) {
        int sum = 0, min = Integer.MAX_VALUE;
        for (int num: nums) {
            sum += num;
            if (num < min) {
                min = num;
            }
        }
        int l = min, r = sum;
        while (l < r) {
            int target = l + (r - l) / 2;
            if (canSplit(nums, m, target)) {
                r = target;
            } else {
                l = target + 1;
            }
        }
        return r;
    }
    
    private boolean canSplit(int[] nums, int m, int target) {
        int count = 1, curSum = 0;
        for (int num: nums) {
            if (num > target) {
                return false;
            }
            curSum += num;
            if (curSum > target) {
                curSum = num;
                count++;
            }
        }
        return count <= m;
    }
}