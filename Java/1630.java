class Solution {
    public List<Boolean> checkArithmeticSubarrays(int[] nums, int[] l, int[] r) {
        List<Boolean> res = new ArrayList<>();
        for (int i = 0; i < l.length; i++) {
            if (r[i] - l[i] < 2) {
                res.add(true);
            } else {
                List<Integer> arr = new ArrayList<>();
                for (int j = l[i]; j <= r[i]; j++) {
                    arr.add(nums[j]);
                }
                Collections.sort(arr);
                int gap = arr.get(1) - arr.get(0);
                boolean arithmetic = true;
                for (int j = 2; j < arr.size(); j++) {
                    if (arr.get(j) - arr.get(j-1) != gap) {
                        arithmetic = false;
                        break;
                    }
                }
                res.add(arithmetic);
            }
        }
        return res;
    }
}