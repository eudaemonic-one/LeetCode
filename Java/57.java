class Solution {
    public int[][] insert(int[][] intervals, int[] newInterval) {
        LinkedList<int[]> res = new LinkedList<>();
        int n = intervals.length;
        int idx = 0;
        int newStart = newInterval[0];
        int newEnd = newInterval[1];
        // append all intervals starting before new interval
        while (idx < n && newStart > intervals[idx][0]) {
            res.add(intervals[idx]);
            idx += 1;
        }
        // if there is no overlap, append new interval
        if (res.isEmpty() || res.getLast()[1] < newStart) {
            res.add(newInterval);
        } else { // otherwise, merge new interval with last one
            int[] interval = res.removeLast();
            if (newEnd > interval[1]) {
                interval[1] = newEnd;
            }
            res.add(interval);
        }
        // append all following intervals
        while (idx < n) {
            int[] interval = intervals[idx];
            int start = interval[0];
            int end = interval[1];
            if (res.getLast()[1] < start) {
                res.add(interval);
            } else {
                interval = res.removeLast();
                if (end > interval[1]) {
                    interval[1] = end;
                }
                res.add(interval);
            }
            idx += 1;
        }
        return res.toArray(new int[res.size()][2]);
    }
}