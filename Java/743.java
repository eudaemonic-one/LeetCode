class Solution {
    public int networkDelayTime(int[][] times, int N, int K) {
        Map<Integer, Integer> distMap = new HashMap<>();
        Map<Integer, List<int[]>> adjList = new HashMap<>();
        for (int[] edge : times) {
            int start = edge[0];
            int end = edge[1];
            int weight = edge[2];
            if (!adjList.containsKey(start)) {
                adjList.put(start, new ArrayList<>());
            }
            adjList.get(start).add(new int[]{end, weight});
        }
        PriorityQueue<int[]> heap = new PriorityQueue<int[]>((a, b) -> a[1] - b[1]);
        heap.offer(new int[]{K, 0});
        while (!heap.isEmpty()) {
            int[] pair = heap.poll();
            int node = pair[0];
            int depth = pair[1];
            if (distMap.containsKey(node)) {
                continue;
            }
            distMap.put(node, depth);
            if (adjList.containsKey(node)) {
                for (int[] edge: adjList.get(node)) {
                    int nei = edge[0];
                    int dist = edge[1];
                    if (!distMap.containsKey(nei)) {
                        heap.offer(new int[]{nei, depth+dist});
                    }
                }
            }
        }
        if (distMap.size() != N) {
            return -1;
        }
        int res = 0;
        for (int dist : distMap.values()) {
            res = Math.max(res, dist);
        }
        return res;
    }
}

// Approach 1: Depth-first Search
// class Solution {
//     public int networkDelayTime(int[][] times, int N, int K) {
//         Map<Integer, Integer> distMap = new HashMap<>();
//         Map<Integer, List<int[]>> adjList = new HashMap<>();
//         for (int[] edge : times) {
//             int start = edge[0];
//             int end = edge[1];
//             int weight = edge[2];
//             if (!adjList.containsKey(start)) {
//                 adjList.put(start, new ArrayList<>());
//             }
//             adjList.get(start).add(new int[]{end, weight});
//         }
//         for (int node : adjList.keySet()) {
//             Collections.sort(adjList.get(node), (a, b) -> a[1] - b[1]);
//         }
//         for (int i = 1; i <= N; i++) {
//             distMap.put(i, Integer.MAX_VALUE);
//         }
//         dfs(adjList, distMap, K, 0);
//         int res = 0;
//         for (int dist : distMap.values()) {
//             if (dist == Integer.MAX_VALUE) {
//                 return -1;
//             }
//             res = Math.max(res, dist);
//         }
//         return res;
//     }
    
//     private void dfs(Map<Integer, List<int[]>> adjList, Map<Integer, Integer> distMap, int node, int depth) {
//         if (depth >= distMap.get(node)) {
//             return;
//         }
//         distMap.put(node, depth);
//         if (adjList.containsKey(node)) {
//             for (int[] pair : adjList.get(node)) {
//                 dfs(adjList, distMap, pair[0], depth+pair[1]);
//             }
//         }
//     }
// }