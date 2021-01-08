class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> (b - a));
        for (int weight: stones) {
            pq.offer(weight);
        }
        while (pq.size() >= 2) {
            int y = pq.poll();
            int x = pq.poll();
            if (x != y) {
                pq.offer(y-x);
            }
        }
        if (pq.isEmpty()) {
            return 0;
        }
        return pq.poll();
    }
}