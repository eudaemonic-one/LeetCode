class Solution {
    public boolean isPossible(int[] nums) {
        Map<Integer, Integer> counter = new HashMap<>();
        Map<Integer, Integer> chains = new HashMap<>();
        for (int x : nums) {
            counter.put(x, counter.getOrDefault(x, 0) + 1);
        }
        for (int x : nums) {
            if (counter.get(x) == 0) {
                continue;
            } else if (chains.getOrDefault(x, 0) > 0) {
                chains.put(x, chains.getOrDefault(x, 0) - 1);
                chains.put(x+1, chains.getOrDefault(x+1, 0) + 1);
            } else if (counter.getOrDefault(x+1, 0) > 0 && counter.getOrDefault(x+2, 0) > 0) {
                counter.put(x+1, counter.getOrDefault(x+1, 0) - 1);
                counter.put(x+2, counter.getOrDefault(x+2, 0) - 1);
                chains.put(x+3, chains.getOrDefault(x+3, 0) + 1);
            } else {
                return false;
            }
            counter.put(x, counter.getOrDefault(x, 0) - 1);
        }
        return true;
    }
}