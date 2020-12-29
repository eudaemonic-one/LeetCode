class Solution {
    public List<Integer> powerfulIntegers(int x, int y, int bound) {
        Set<Integer> seen = new HashSet<>();
        for (int i = 0; i <= 19; i++) {
            int powX = (int) Math.pow(x, i);
            if (powX > bound) {
                break;
            }
            for (int j = 0; j <= 19; j++) {
                int powY = (int) Math.pow(y, j);
                if (powY > bound) {
                    break;
                }
                int val = powX + powY;
                if (val <= bound) {
                    seen.add(val);
                }
            }
        }
        return new ArrayList(seen);
    }
}