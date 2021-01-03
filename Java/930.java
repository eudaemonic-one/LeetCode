// Approach 2: Three Pointer
class Solution {
    public int numSubarraysWithSum(int[] A, int S) {
        int res = 0;
        int lo = 0, hi = 0;
        int sumLo = 0, sumHi = 0;
        for (int j = 0; j < A.length; j++) {
            sumLo += A[j];
            while (lo < j && sumLo > S) {
                sumLo -= A[lo++];
            }
            sumHi += A[j];
            while (hi < j && (sumHi > S || sumHi == S && A[hi] == 0)) {
                sumHi -= A[hi++];
            }
            if (sumLo == S) {
                res += hi - lo + 1;
            }
        }
        return res;
    }
}

// Approach 1: Prefix Sums
// class Solution {
//     public int numSubarraysWithSum(int[] A, int S) {
//         int n = A.length;
//         int[] presum = new int[n+1];
//         for (int i = 0; i < n; i++) {
//             presum[i+1] = presum[i] + A[i];
//         }
//         int res = 0;
//         Map<Integer, Integer> counter = new HashMap<>();
//         for (int x : presum) {
//             res += counter.getOrDefault(x, 0);
//             counter.put(x+S, counter.getOrDefault(x+S, 0) + 1);
//         }
//         return res;
//     }
// }