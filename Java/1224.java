class Solution {
    public int maxEqualFreq(int[] nums) {
        int[] freq = new int[100001];
        int[] count = new int[100001];
        int res = 0;
        for (int i = 1; i <= nums.length; i++) {
            int num = nums[i-1];
            freq[count[num]] -= 1;
            count[num] += 1;
            freq[count[num]] += 1;
            if (freq[count[num]] * count[num] == i && i < nums.length) {
                res = i + 1;
            }
            int rem = i - freq[count[num]] * count[num];
            if ((rem == 1 || rem == count[num] + 1) && freq[rem] == 1) {
                res = i;
            }
        }
        return res;
    }
}