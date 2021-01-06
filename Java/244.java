class WordDistance {
    private Map<String, List<Integer>> distMap;
    
    public WordDistance(String[] words) {
        distMap = new HashMap<>();
        for (int i = 0; i < words.length; i++) {
            List<Integer> index = distMap.getOrDefault(words[i], new ArrayList<>());
            index.add(i);
            distMap.put(words[i], index);
        }
    }
    
    public int shortest(String word1, String word2) {
        List<Integer> index1 = distMap.get(word1);
        List<Integer> index2 = distMap.get(word2);
        int i1 = 0, i2 = 0, minDiff = Integer.MAX_VALUE;
        while (i1 < index1.size() && i2 < index2.size()) {
            minDiff = Math.min(minDiff, Math.abs(index1.get(i1) - index2.get(i2)));
            if (index1.get(i1) < index2.get(i2)) {
                i1++;
            } else {
                i2++;
            }
        }
        return minDiff;
    }
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * WordDistance obj = new WordDistance(words);
 * int param_1 = obj.shortest(word1,word2);
 */