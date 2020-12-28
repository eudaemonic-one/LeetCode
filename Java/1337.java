// Approach 3: Binary Search and Priority Queue
class Solution {
    public int[] kWeakestRows(int[][] mat, int k) {
        int m = mat.length;
        int n = mat[0].length;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> {
            if (a[1] == b[1]) {
                return b[0] - a[0];
            }
            return b[1] - a[1];
        });
        for (int i = 0; i < m; i++) {
            int count = binarySearch(mat[i]);
            int[] pair = new int[]{i, count};
            pq.add(pair);
            if (pq.size() > k) {
                pq.poll();
            }
        }
        int[] res = new int[k];
        for (int i = k-1; i >= 0; i--) {
            res[i] = pq.poll()[0];
        }
        return res;
    }
    
    private int binarySearch(int[] nums) {
        int l = 0;
        int r = nums.length;
        while (l < r) {
            int m = l + (r - l) / 2;
            if (nums[m] == 1) {
                l = m + 1;
            } else {
                r = m;
            }
        }
        return l;
    }
}

// Approach 2: Binary Search and Sorting/Map
// class Solution {
//     public int[] kWeakestRows(int[][] mat, int k) {
//         int m = mat.length;
//         int n = mat[0].length;
//         Map<Integer, List<Integer>> strengthMap = new TreeMap<>();
//         for (int i = 0; i < m; i++) {
//             int count = binarySearch(mat[i]);
//             if (!strengthMap.containsKey(count)) {
//                 strengthMap.put(count, new ArrayList<>());
//             }
//             strengthMap.get(count).add(i);
//         }
//         int[] res = new int[k];
//         int i = 0;
//         for (int key : strengthMap.keySet()) {
//             for (int index : strengthMap.get(key)) {
//                 res[i] = index;
//                 ++i;
//                 if (i >= k) {
//                     return res;
//                 }
//             }
//         }
//         return res;
//     }
    
//     private int binarySearch(int[] nums) {
//         int l = 0;
//         int r = nums.length;
//         while (l < r) {
//             int m = l + (r - l) / 2;
//             if (nums[m] == 1) {
//                 l = m + 1;
//             } else {
//                 r = m;
//             }
//         }
//         return l;
//     }
// }

// Approach 1: Linear Serach and Sorting
// class Solution {
//     public int[] kWeakestRows(int[][] mat, int k) {
//         int m = mat.length;
//         int n = mat[0].length;
//         int[][] pairs = new int[m][2];
//         for (int i = 0; i < m; i++) {
//             int count = 0;
//             for (int j = 0; j < n; j++) {
//                 if (mat[i][j] == 1) {
//                     count += 1;
//                 }
//             }
//             pairs[i][0] = i;
//             pairs[i][1] = count;
//         }
//         Arrays.sort(pairs, (a, b) -> {
//             if (a[1] == b[1]) {
//                 return a[0] - b[0];
//             }
//             return a[1] - b[1];
//         });
//         int[] res = new int[k];
//         for (int i = 0; i < k; i++) {
//             res[i] = pairs[i][0];
//         }
//         return res;
//     }
// }