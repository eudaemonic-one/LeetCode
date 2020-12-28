class Solution {
    private enum State {
        UNVISITED,
        VISITING,
        VISITED;
    }
    
    public String alienOrder(String[] words) {
        Map<Character, Integer> counter = new HashMap<>();
        Map<Character, List<Character>> adjList = new HashMap<>();
        for (String word : words) {
            for (char ch : word.toCharArray()) {
                counter.put(ch, 0);
                adjList.put(ch, new ArrayList<>());
            }
        }
        for (int i = 0; i < words.length-1; i++) {
            String word1 = words[i];
            String word2 = words[i+1];
            if (word1.length() > word2.length() && word1.startsWith(word2)) {
                return "";
            }
            for (int j = 0; j < Math.min(word1.length(), word2.length()); j++) {
                char ch1 = word1.charAt(j);
                char ch2 = word2.charAt(j);
                if (ch1 != ch2) {
                    adjList.get(ch1).add(ch2);
                    counter.put(ch2, counter.get(ch2) + 1);
                    break;
                }
            }
        }
        StringBuilder sb = new StringBuilder();
        Queue<Character> queue = new LinkedList<>();
        for (Map.Entry<Character, Integer> entry : counter.entrySet()) {
            if (entry.getValue() == 0) {
                queue.add(entry.getKey());
            }
        }
        while (!queue.isEmpty()) {
            char ch = queue.poll();
            sb.append(ch);
            for (char nei : adjList.get(ch)) {
                counter.put(nei, counter.get(nei) - 1);
                if (counter.get(nei) == 0) {
                    queue.add(nei);
                }
            }
        }
        if (sb.length() < counter.size()) {
            return "";
        }
        return sb.toString();
    }
}