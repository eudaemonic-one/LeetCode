class Solution {
    public List<Integer> largestDivisibleSubset(int[] nums) {
        int n = nums.length;
        if (n == 0) {
            return new ArrayList<>();
        }
        List<List<Integer>> cand = new ArrayList<>();
        for (int num: nums) {
            cand.add(new ArrayList<>());
        }
        Arrays.sort(nums);
        for (int i = 0; i < n; i++) {
            List<Integer> maxSubset = new ArrayList();
            for (int k = 0; k < i; k++) {
                if (nums[i] % nums[k] == 0 && cand.get(k).size() > maxSubset.size()) {
                    maxSubset = cand.get(k);
                }
            }
            cand.get(i).addAll(maxSubset);
            cand.get(i).add(nums[i]);
        }
        List<Integer> res = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (cand.get(i).size() > res.size()) {
                res = cand.get(i);
            }
        }
        return res;
    }
}