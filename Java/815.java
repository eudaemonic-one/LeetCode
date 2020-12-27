class Solution {    
    public int numBusesToDestination(int[][] routes, int S, int T) {
        if (S == T) {
            return 0;
        }
        Map<Integer, Set<Integer>> routeMap = new HashMap<>();
        for (int i = 0; i < routes.length; i++) {
            for (int v : routes[i]) {
                if (!routeMap.containsKey(v)) {
                    routeMap.put(v, new HashSet<>());
                }
                routeMap.get(v).add(i);
            }
        }
        Queue<int[]> queue = new ArrayDeque<>();
        Set<Integer> visitedStops = new HashSet<>();
        Set<Integer> visitedRoutes = new HashSet<>();
        queue.offer(new int[]{S, 0});
        visitedStops.add(S);
        while (!queue.isEmpty()) {
            int[] pair = queue.poll();
            int stop = pair[0];
            int depth = pair[1];
            if (stop == T) {
                return depth;
            }
            for (int route : routeMap.get(stop)) {
                if (visitedRoutes.contains(route)) {
                    continue;
                }
                for (int s : routes[route]) {
                    if (!visitedStops.contains(s)) {
                        visitedStops.add(s);
                        queue.offer(new int[]{s, depth+1});
                    }
                }
                visitedRoutes.add(route);
            }
        }
        return -1;
    }
}
