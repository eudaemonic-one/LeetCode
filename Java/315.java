class Solution {
    public List<Integer> countSmaller(int[] nums) {
        List<Integer> res = new ArrayList<>();
        int n = nums.length;
        int[] indexes = new int[n];
        for (int i = 0; i < n; i++) {
            indexes[i] = i;
        }
        int[] count = new int[n];
        mergeSort(0, n - 1, count, nums, indexes);
        for (int i = 0; i < n; i++) {
            res.add(count[i]);
        }
        return res;
    }
    
    private void mergeSort(int l, int r, int[] count, int[] nums, int[] indexes) {
        if (l >= r) {
            return;
        }
        int m = l + (r - l) / 2;
        // split the array
        mergeSort(l, m, count, nums, indexes);
        mergeSort(m + 1, r, count, nums, indexes);
        // merge two subarrays
        int swap = 0;
        int[] newIndexes = new int[r - l + 1];
        int pL = l, pR = m + 1, idx = 0;
        while (pL <= m && pR <= r) {
            if (nums[indexes[pL]] > nums[indexes[pR]]) {
                newIndexes[idx] = indexes[pR];
                pR++;
                swap++;
            } else {
                newIndexes[idx] = indexes[pL];
                count[indexes[pL]] += swap;
                pL++;
            }
            idx++;
        }
        while (pL <= m) {
            newIndexes[idx] = indexes[pL];
            count[indexes[pL]] += swap;
            pL++;
            idx++;
        }
        while (pR <= r) {
            newIndexes[idx] = indexes[pR];
            pR++;
            idx++;
        }
        for (int i = l; i <= r; i++) {
            indexes[i] = newIndexes[i - l];
        }
    }
}