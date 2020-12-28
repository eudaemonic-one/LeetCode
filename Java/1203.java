class Solution {
    public int[] sortItems(int n, int m, int[] group, List<List<Integer>> beforeItems) {
        Map<Integer, List<Integer>> itemAdjList = new HashMap<>();
        for (int i = 0; i < n; i++) {
            itemAdjList.put(i, new ArrayList<>());
        }
        int[] itemIndegrees = new int[n];
        for (int toItem = 0; toItem < n; toItem++) {
            List<Integer> fromItems = beforeItems.get(toItem);
            for (int fromItem : fromItems) {
                itemAdjList.get(fromItem).add(toItem);
                itemIndegrees[toItem] += 1;
            }
        }
        
        Map<Integer, List<Integer>> groupAdjList = new HashMap<>();
        for (int i = 0; i < group.length; i++) {
            if (group[i] == -1) {
                group[i] = m++;
            }
        }
        for (int i = 0; i < m; i++) {
            groupAdjList.put(i, new ArrayList<>());
        }
        int[] groupIndegrees = new int[m];        
        for (int i = 0; i < group.length; i++) {
            int toGroup = group[i];
            List<Integer> fromItems = beforeItems.get(i);
            for (int fromItem : fromItems) {
                int fromGroup = group[fromItem];
                if (fromGroup != toGroup) {
                    groupAdjList.get(fromGroup).add(toGroup);
                    groupIndegrees[toGroup] += 1;
                }
            }
        }
        
        List<Integer> itemOrder = bfs(itemAdjList, n, itemIndegrees);
        List<Integer> groupOrder = bfs(groupAdjList, m, groupIndegrees);
        
        if (itemOrder.isEmpty() || groupOrder.isEmpty()) {
            return new int[0];
        }
        Map<Integer, List<Integer>> groupToItems = new HashMap<>();
        for (int item : itemOrder) {
            if (!groupToItems.containsKey(group[item])) {
                groupToItems.put(group[item], new ArrayList<>());
            }
            groupToItems.get(group[item]).add(item);
        }
        int[] res = new int[n];
        int idx = 0;
        for (int groupId : groupOrder) {
            for (int item : groupToItems.getOrDefault(groupId, new ArrayList<>())) {
                res[idx++] = item;
            }
        }
        return res;
    }
    
    private List<Integer> bfs(Map<Integer, List<Integer>> adjList, int n, int[] indegrees) {
        Queue<Integer> queue = new LinkedList<>();
        for (int key : adjList.keySet()) {
            if (indegrees[key] == 0) {
                queue.add(key);
            }
        }
        List<Integer> order = new ArrayList<>();
        while (!queue.isEmpty()) {
            int node = queue.poll();
            order.add(node);
            for (int nei : adjList.get(node)) {
                indegrees[nei] -= 1;
                if (indegrees[nei] == 0) {
                    queue.add(nei);
                }
            }
        }
        if (order.size() < n) {
            return new ArrayList<>();
        }
        return order;
    }
}