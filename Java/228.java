// Approach 2: Simulation (Alternative)
class Solution {
    public List<String> summaryRanges(int[] nums) {
        List<String> res = new ArrayList<>();
        for (int i = 0, j = 0; j < nums.length; j++) {
            if (j + 1 < nums.length && nums[j+1] == nums[j] + 1) {
                continue;
            }
            if (i == j) {
                res.add(nums[i] + "");
            } else {
                res.add(nums[i] + "->" + nums[j]);
            }
            i = j + 1;
        }
        return res;
    }
}

// Approach 1: Simulation
// class Solution {
//     public List<String> summaryRanges(int[] nums) {
//         List<String> res = new ArrayList<>();
//         if (nums.length == 0) {
//             return res;
//         }
//         int l = nums[0], r = nums[0], lastL = Integer.MIN_VALUE;
//         for (int i = 1; i < nums.length; i++) {
//             if (nums[i] > r + 1) {
//                 if (l == r) {
//                     res.add(String.format("%d", l));
//                 } else {
//                     res.add(String.format("%d->%d", l, r));
//                 }
//                 lastL = l;
//                 l = nums[i];
//                 r = nums[i];
//             } else {
//                 r = nums[i];
//             }
//         }
//         if (lastL != nums[nums.length-1]) {
//             if (l == r) {
//                 res.add(String.format("%d", l));
//             } else {
//                 res.add(String.format("%d->%d", l, r));
//             }
//         }
//         return res;
//     }
// }