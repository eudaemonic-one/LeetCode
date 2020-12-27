class Solution {
    private enum State {
        UNVISITED,
        VISITING,
        VISITED;
    }
    
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        boolean isValid = true;
        Map<Integer, State> stateTable = new HashMap<>();
        Map<Integer, List<Integer>> adjList = new HashMap<>();
        List<Integer> order = new ArrayList<>();
        // initialize empty states
        for (int i = 0; i < numCourses; i++) {
            stateTable.put(i, State.UNVISITED);
        }
        // build adjacent list for the topological graph
        for (int[] path : prerequisites) {
            int dst = path[0];
            int src = path[1];
            List<Integer> neighbors = adjList.getOrDefault(src, new ArrayList());
            neighbors.add(dst);
            adjList.put(src, neighbors);
        }
        // dfs every unprocessed node
        for (int i = 0; i < numCourses; i++) {
            if (stateTable.get(i) == State.UNVISITED) {
                if (!dfs(i, stateTable, adjList, order)) {
                    return new int[0];
                }
            }
        }
        // reconstruct solution
        int[] res = new int[numCourses];
        for (int i = 0; i < numCourses; i++) {
            res[i] = order.get(numCourses - i - 1);
        }
        return res;
    }
    
    private boolean dfs(int v, Map<Integer, State> stateTable, Map<Integer, List<Integer>> adjList, List<Integer> order) {
        stateTable.put(v, State.VISITING); // start the recursion
        for (int nei : adjList.getOrDefault(v, new ArrayList<>())) {
            if (stateTable.get(nei) == State.UNVISITED) {
                if (!dfs(nei, stateTable, adjList, order)) {
                    return false; // invalid
                }
            } else if (stateTable.get(nei) == State.VISITING) {
                return false; // exists cycle
            }
        }
        stateTable.put(v, State.VISITED); // mark as visited
        order.add(v);
        return true;
    }
}