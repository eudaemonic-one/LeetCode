class Solution {
    public List<List<Integer>> getSkyline(int[][] buildings) {
        List<List<Integer>> res = new ArrayList<>();
        int n = buildings.length;
        if (n == 0) {
            return res;
        }
        if (n == 1) {
            int xStart = buildings[0][0];
            int xEnd = buildings[0][1];
            int y = buildings[0][2];
            return new ArrayList<>(List.of(List.of(xStart, y), List.of(xEnd, 0)));
        }
        List<List<Integer>> leftSkyline = getSkyline(Arrays.copyOfRange(buildings, 0, n/2));
        List<List<Integer>> rightSkyline = getSkyline(Arrays.copyOfRange(buildings, n/2, n));
        return mergeSkylines(leftSkyline, rightSkyline);
    }
    
    private List<List<Integer>> mergeSkylines(List<List<Integer>> left, List<List<Integer>> right) {
        int nL = left.size(), nR = right.size();
        int l = 0, r = 0;
        int currY = 0, leftY = 0, rightY = 0;
        List<List<Integer>> res = new ArrayList<>();
        while (l < nL && r < nR) {
            List<Integer> pointL = left.get(l);
            List<Integer> pointR = right.get(r);
            int x;
            if (pointL.get(0) < pointR.get(0)) {
                x = pointL.get(0);
                leftY = pointL.get(1);
                l++;
            } else {
                x = pointR.get(0);
                rightY = pointR.get(1);
                r++;
            }
            int maxY = Math.max(leftY, rightY);
            if (currY != maxY) {
                updateSkyline(res, x, maxY);
                currY = maxY;
            }
        }
        appendSkyline(res, left, l, nL, currY);
        appendSkyline(res, right, r, nR, currY);
        return res;
    }
    
    private void updateSkyline(List<List<Integer>> res, int x, int y) {
        if (res.isEmpty() || res.get(res.size()-1).get(0) != x) {
            res.add(new ArrayList<>(List.of(x, y)));
        } else {
            res.get(res.size()-1).set(1, y);
        }
    }
    
    private void appendSkyline(List<List<Integer>> res, List<List<Integer>> skyline,
                            int i, int n, int currY) {
        while (i < n) {
            List<Integer> point = skyline.get(i);
            int x = point.get(0);
            int y = point.get(1);
            if (y != currY) {
                updateSkyline(res, x, y);
                currY = y;
            }
            i++;
        }
    }
}