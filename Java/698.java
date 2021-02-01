class Solution {
    public boolean canPartitionKSubsets(int[] nums, int k) {
        Arrays.sort(nums);
        int sum = 0;
        for (int num: nums) {
            sum += num;
        }
        if (sum % k != 0) {
            return false;
        }
        int target = sum / k;
        if (nums[nums.length-1] > target) {
            return false;
        }
        int idx = nums.length - 1;
        while (idx >= 0 && nums[idx] == target) {
            idx--;
            k--;
        }
        int[] groups = new int[k];
        return dfs(nums, k, target, idx, groups);
    }
    
    private boolean dfs(int[] nums, int k, int target, int idx, int[] groups) {
        if (idx < 0) {
            return true;
        }
        int val = nums[idx];
        for (int i = 0; i < k; i++) {
            if (groups[i] + val <= target) {
                groups[i] += val;
                if (dfs(nums, k, target, idx-1, groups)) {
                    return true;
                }
                groups[i] -= val;
            }
            if (groups[i] == 0) {
                break;
            }
        }
        return false;
    }
}