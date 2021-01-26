class Solution {
    public int findMinMoves(int[] machines) {
        int n = machines.length;
        int dressTotal = 0;
        for (int dress : machines) {
            dressTotal += dress;
        }
        if (dressTotal % n != 0) {
            return -1;
        }
        int dressPerMachine = dressTotal / n;
        for (int i = 0; i < n; i++) {
            machines[i] -= dressPerMachine;
        }
        int res = 0, currSum = 0, maxSum = 0, tmpRes = 0;
        for (int dress : machines) {
            currSum += dress;
            maxSum = Math.max(maxSum, Math.abs(currSum));
            res = Math.max(res, Math.max(maxSum, dress));
        }
        return res;
    }
}