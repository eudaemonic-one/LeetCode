class Solution {
    public List<String> findMissingRanges(int[] nums, int lower, int upper) {
        List<String> missingRanges = new ArrayList<>();
        int n = nums.length;
        if (n == 0) {
            missingRanges.add(formatRange(lower, upper));
            return missingRanges;
        }
        if (nums[0] > lower) {
            missingRanges.add(formatRange(lower, nums[0] - 1));
        }
        for (int i = 1; i < n; i++) {
            if (nums[i] - nums[i-1] > 1) {
                missingRanges.add(formatRange(nums[i-1] + 1, nums[i] - 1));
            }
        }
        if (nums[n-1] < upper) {
            missingRanges.add(formatRange(nums[n-1] + 1, upper));
        }
        return missingRanges;
    }
    
    private String formatRange(int lower, int upper) {
        if (lower == upper) {
            return String.valueOf(lower);
        } else {
            return lower + "->" + upper;
        }
    }
}