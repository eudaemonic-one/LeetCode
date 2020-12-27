class Solution {
    public List<Boolean> checkArithmeticSubarrays(int[] nums, int[] l, int[] r) {
        List<Boolean> res = new ArrayList<>();
        for (int i = 0; i < l.length; i++) {
            if (r[i] - l[i] < 2) {
                res.add(true);
            } else {
                int len = r[i] - l[i] + 1;
                int min = Integer.MAX_VALUE;
                int max = Integer.MIN_VALUE;
                for (int j = l[i]; j <= r[i]; j++) {
                    if (nums[j] < min) {
                        min = nums[j];
                    }
                    if (nums[j] > max) {
                        max = nums[j];
                    }
                }
                int rem = (max - min) % (len - 1);
                if (min == max) {
                    res.add(true);
                } else if (rem != 0) {
                    res.add(false);
                } else {
                    int gap = (max - min) / (len - 1);
                    boolean[] visited = new boolean[len];
                    boolean arithmetic = true;
                    for (int j = l[i]; j <= r[i]; j++) {
                        int diff = nums[j] - min;
                        if (diff % gap != 0 || visited[diff/gap]) {
                            arithmetic = false;
                            break;
                        }
                        visited[diff/gap] = true;
                    }
                    res.add(arithmetic);
                }
            }
        }
        return res;
    }
}

// Approach 1: Sort
// class Solution {
//     public List<Boolean> checkArithmeticSubarrays(int[] nums, int[] l, int[] r) {
//         List<Boolean> res = new ArrayList<>();
//         for (int i = 0; i < l.length; i++) {
//             if (r[i] - l[i] < 2) {
//                 res.add(true);
//             } else {
//                 List<Integer> arr = new ArrayList<>();
//                 for (int j = l[i]; j <= r[i]; j++) {
//                     arr.add(nums[j]);
//                 }
//                 Collections.sort(arr);
//                 int gap = arr.get(1) - arr.get(0);
//                 boolean arithmetic = true;
//                 for (int j = 2; j < arr.size(); j++) {
//                     if (arr.get(j) - arr.get(j-1) != gap) {
//                         arithmetic = false;
//                         break;
//                     }
//                 }
//                 res.add(arithmetic);
//             }
//         }
//         return res;
//     }
// }