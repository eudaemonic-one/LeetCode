// Approach #4: Using Priority Queue
class Solution {
    public int[] smallestRange(List<List<Integer>> nums) {
        int minx = 0, miny = Integer.MAX_VALUE, max = Integer.MIN_VALUE;
        int[] next = new int[nums.size()];
        boolean flag = true;
        PriorityQueue<Integer> minQueue = new PriorityQueue<>((i, j) -> nums.get(i).get(next[i]) - nums.get(j).get(next[j]));
        for (int i = 0; i < nums.size(); i++) {
            minQueue.offer(i);
            max = Math.max(max, nums.get(i).get(0));
        }
        for (int i = 0; i < nums.size() && flag; i++) {
            for (int j = 0; j < nums.get(i).size() && flag; j++) {
                int mini = minQueue.poll();
                if (miny - minx > max - nums.get(mini).get(next[mini])) {
                    minx = nums.get(mini).get(next[mini]);
                    miny = max;
                }
                next[mini]++;
                if (next[mini] == nums.get(mini).size()) {
                    flag = false;
                    break;
                }
                minQueue.offer(mini);
                max = Math.max(max, nums.get(mini).get(next[mini]));
            }
        }
        return new int[]{minx, miny};
    }
}

// Approach #1: Using Pointers
// class Solution {
//     public int[] smallestRange(List<List<Integer>> nums) {
//         int minx = 0, miny = Integer.MAX_VALUE;
//         int[] next = new int[nums.size()];
//         boolean flag = true;
//         for (int i = 0; i < nums.size() && flag; i++) {
//             for (int j = 0; j < nums.get(i).size() && flag; j++) {
//                 int mini = 0, maxi = 0;
//                 for (int k = 0; k < nums.size(); k++) {
//                     if (nums.get(mini).get(next[mini]) > nums.get(k).get(next[k])) {
//                         mini = k;
//                     }
//                     if (nums.get(maxi).get(next[maxi]) < nums.get(k).get(next[k])) {
//                         maxi = k;
//                     }
//                 }
//                 if (miny - minx > nums.get(maxi).get(next[maxi]) - nums.get(mini).get(next[mini])) {
//                     minx = nums.get(mini).get(next[mini]);
//                     miny = nums.get(maxi).get(next[maxi]);
//                 }
//                 next[mini]++;
//                 if (next[mini] == nums.get(mini).size()) {
//                     flag = false;
//                 }
//             }
//         }
//         return new int[]{minx, miny};
//     }
// }