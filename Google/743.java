// Approach #2: Dijkstra's Algorithm

class Solution {
    public int networkDelayTime(int[][] times, int n, int k) {
        Map<Integer, List<int[]>> graph = new HashMap<>();
        for (int[] edge : times) {
            if (!graph.containsKey(edge[0])) {
                graph.put(edge[0], new ArrayList<>());
            }
            graph.get(edge[0]).add(new int[]{edge[1], edge[2]});
        }
        
        PriorityQueue<int[]> heap = new PriorityQueue<>((pair1, pair2) -> pair1[1] - pair2[1]);
        heap.offer(new int[]{k, 0});
        Map<Integer, Integer> distMap = new HashMap<>();
        
        while (!heap.isEmpty()) {
            int[] pair = heap.poll();
            int node = pair[0], elapsed = pair[1];
            if (distMap.containsKey(node)) {
                continue;
            }
            distMap.put(node, elapsed);
            if (graph.containsKey(node)) {
                for (int[] edge : graph.get(node)) {
                    int nei = edge[0], dist = edge[1];
                    if (!distMap.containsKey(nei)) {
                        heap.offer(new int[]{nei, elapsed+dist});
                    }
                }
            }
        }
        
        if (distMap.size() != n) {
            return -1;
        }
        int res = 0;
        for (int cand : distMap.values()) {
            res = Math.max(res, cand);
        }
        return res;
    }
}

// Approach #1: Depth-First Search

// class Solution {
//     public int networkDelayTime(int[][] times, int n, int k) {
//         Map<Integer, List<int[]>> graph = new HashMap<>();
//         for (int[] edge : times) {
//             if (!graph.containsKey(edge[0])) {
//                 graph.put(edge[0], new ArrayList<>());
//             }
//             graph.get(edge[0]).add(new int[]{edge[1], edge[2]});
//         }
//         for (int node: graph.keySet()) {
//             Collections.sort(graph.get(node), (a, b) -> a[1] - b[1]);
//         }
        
//         Map<Integer, Integer> distMap = new HashMap<>();
//         for (int node = 1; node <= n; node++) {
//             distMap.put(node, Integer.MAX_VALUE);
//         }
        
//         dfs(graph, distMap, k, 0);
        
//         int res = 0;
//         for (int cand : distMap.values()) {
//             if (cand == Integer.MAX_VALUE) {
//                 return -1;
//             }
//             res = Math.max(res, cand);
//         }
//         return res;
//     }
    
//     private void dfs(Map<Integer, List<int[]>> graph, Map<Integer, Integer> distMap, int node, int elapsed) {
//         if (elapsed >= distMap.get(node)) {
//             return;
//         }
//         distMap.put(node, elapsed);
//         if (graph.containsKey(node)) {
//             for (int[] pair : graph.get(node)) {
//                 dfs(graph, distMap, pair[0], elapsed+pair[1]);
//             }
//         }
//     }
// }