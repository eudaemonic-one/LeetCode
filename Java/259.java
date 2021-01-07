class Solution {
    public int threeSumSmaller(int[] nums, int target) {
        Arrays.sort(nums);
        int res = 0;
        for (int i = 0; i < nums.length-2; i++) {
            res += twoSumSmaller(nums, i+1, target-nums[i]);
        }
        return res;
    }
    
    private int twoSumSmaller(int[] nums, int startIdx, int target) {
        int sum = 0, j = startIdx, k = nums.length - 1;
        while (j < k) {
            if (nums[j] + nums[k] < target) {
                sum += k - j;
                j++;
            } else {
                k--;
            }
        }
        return sum;
    }
}