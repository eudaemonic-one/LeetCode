class Solution {
    public int maximumRequests(int n, int[][] requests) {
        int[] res = new int[1];
        backtrack(n, requests, new int[n], 0, 0, res);
        return res[0];
    }
    
    private void backtrack(int n, int[][] requests, int[] counter, int idx, int num, int[] res) {
        if (idx == requests.length) {
            for (int count: counter) {
                if (count != 0) {
                    return;
                }
            }
            res[0] = Math.max(res[0], num);
            return;
        }
        counter[requests[idx][0]]--;
        counter[requests[idx][1]]++;
        backtrack(n, requests, counter, idx+1, num+1, res);
        counter[requests[idx][0]]++;
        counter[requests[idx][1]]--;
        backtrack(n, requests, counter, idx+1, num, res);
    }
}