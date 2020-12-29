class Solution {
    public boolean isMajorityElement(int[] nums, int target) {
        int l = 0;
        int r = nums.length;
        while (l < r) {
            int m = l + (r - l) / 2;
            if (nums[m] < target) {
                l = m + 1;
            } else {
                r = m;
            }
        }
        int startIdx = l;
        int endIdx = l + nums.length / 2;
        return endIdx < nums.length && nums[endIdx] == target;
    }
}