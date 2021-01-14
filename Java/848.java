class Solution {
    public String shiftingLetters(String S, int[] shifts) {
        StringBuilder sb = new StringBuilder();
        int sum = 0;
        for (int shift: shifts) {
            sum = (sum + shift) % 26;
        }
        for (int i = 0; i < S.length(); i++) {
            int index = S.charAt(i) - 'a';
            sb.append((char) ((index + sum) % 26 + 97));
            sum = Math.floorMod(sum - shifts[i], 26);
        }
        return sb.toString();
    }
}