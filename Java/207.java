class Solution {
    private enum State {
        UNVISITED,
        VISITING,
        VISITED;
    }
    
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        Map<Integer, List<Integer>> adjList = new HashMap<>();
        for (int[] path : prerequisites) {
            int dst = path[0];
            int src = path[1];
            List<Integer> neighbors = adjList.getOrDefault(src, new ArrayList<>());
            neighbors.add(dst);
            adjList.put(src, neighbors);
        }
        Map<Integer, State> stateTable = new HashMap<>();
        for (int i = 0; i < numCourses; i++) {
            stateTable.put(i, State.UNVISITED);
        }
        for (int i = 0; i < numCourses; i++) {
            if (stateTable.get(i).equals(State.UNVISITED)) {
                if (!dfs(i, adjList, stateTable)) {
                    return false;
                }
            }
        }
        return true;
    }
    
    private boolean dfs(int v, Map<Integer, List<Integer>> adjList, Map<Integer, State> stateTable) {
        stateTable.put(v, State.VISITING);
        for (int nei : adjList.getOrDefault(v, new ArrayList<>())) {
            if (stateTable.get(nei).equals(State.UNVISITED)) {
                if (!dfs(nei, adjList, stateTable)) {
                    return false;
                }
            } else if (stateTable.get(nei).equals(State.VISITING)) {
                return false;
            }
        }
        stateTable.put(v, State.VISITED);
        return true;
    }
}