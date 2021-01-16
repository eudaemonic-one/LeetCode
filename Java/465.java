class Solution {
    public int minTransfers(int[][] transactions) {
        Map<Integer, Integer> balance = new HashMap<>();
        for (int[] tx: transactions) {
            balance.put(tx[0], balance.getOrDefault(tx[0], 0) - tx[2]);
            balance.put(tx[1], balance.getOrDefault(tx[1], 0) + tx[2]);
        }
        List<Integer> debt = new ArrayList<>();
        for (int bal: balance.values()) {
            if (bal != 0) {
                debt.add(bal);
            }
        }
        return dfs(debt, 0);
    }
    
    private int dfs(List<Integer> debt, int start) {
        while (start < debt.size() && debt.get(start) == 0) {
            start++;
        }
        int res = Integer.MAX_VALUE;
        for (int i = start+1, prev = 0; i < debt.size(); i++) {
            if (debt.get(i) != prev && debt.get(i) * debt.get(start) < 0) {
                prev = debt.get(i);
                debt.set(i, debt.get(i) + debt.get(start));
                res = Math.min(res, 1 + dfs(debt, start+1));
                debt.set(i, prev);
            }
        }
        return res < Integer.MAX_VALUE ? res : 0;
    }
}