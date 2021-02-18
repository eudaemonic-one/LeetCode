class Solution {
    public List<String> wordSubsets(String[] A, String[] B) {
        int[] bmax = new int[26];
        for (String b : B) {
            int[] bCount = count(b);
            for (int i = 0; i < 26; i++) {
                bmax[i] = Math.max(bmax[i], bCount[i]);
            }
        }
        List<String> res = new ArrayList<>();
        for (String a : A) {
            int[] aCount = count(a);
            boolean flag = true;
            for (int i = 0; i < 26; i++) {
                if (aCount[i] < bmax[i]) {
                    flag = false;
                    break;
                }
            }
            if (flag) {
                res.add(a);
            }
        }
        return res;
    }
    
    private int[] count(String word) {
        int[] res = new int[26];
        for (char ch : word.toCharArray()) {
            ++res[ch - 'a'];
        }
        return res;
    }
}