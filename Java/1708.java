class Solution {
    public int[] largestSubarray(int[] nums, int k) {
        int max = 0;
        for(int i = 1; i < nums.length-k+1; i++){
            if(nums[i] > nums[max]) {
                max = i;
            }
        }
        int[] res = new int[k];
        for(int i = 0; i < k; i++) {
            res[i] = nums[max+i];
        }
        return res;
    }
}