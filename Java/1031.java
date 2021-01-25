class Solution {
    public int maxSumTwoNoOverlap(int[] A, int L, int M) {
        for (int i = 1; i < A.length; i++) {
            A[i] += A[i-1];
        }
        int res = A[L + M - 1];
        int LMax = A[L-1], MMax = A[M-1];
        for (int i = L+M; i < A.length; i++) {
            LMax = Math.max(LMax, A[i-M] - A[i-M-L]);
            MMax = Math.max(MMax, A[i-L] - A[i-M-L]);
            res = Math.max(res, Math.max(LMax+A[i]-A[i-M], MMax+A[i]-A[i-L]));
        }
        return res;
    }
}