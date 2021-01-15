class FirstUnique {
    private Set<Integer> uniqueSet = new LinkedHashSet<>();
    private Map<Integer, Boolean> uniqueMap = new HashMap<>();

    public FirstUnique(int[] nums) {
        for (int num: nums) {
            add(num);
        }
    }
    
    public int showFirstUnique() {
        if (uniqueSet.isEmpty()) {
            return -1;
        }
        return uniqueSet.iterator().next();
    }
    
    public void add(int value) {
        if (!uniqueMap.containsKey(value)) {
            uniqueMap.put(value, true);
            uniqueSet.add(value);
        } else if (uniqueMap.get(value)) {
            uniqueMap.put(value, false);
            uniqueSet.remove(value);
        }
    }
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * FirstUnique obj = new FirstUnique(nums);
 * int param_1 = obj.showFirstUnique();
 * obj.add(value);
 */