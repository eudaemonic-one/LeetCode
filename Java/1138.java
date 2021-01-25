class Solution {
    public String alphabetBoardPath(String target) {
        int r = 0, c = 0;
        StringBuilder sb = new StringBuilder();
        for (char ch : target.toCharArray()) {
            int r1 = (ch - 'a') / 5;
            int c1 = (ch - 'a') % 5;
            sb.append(String.join("", Collections.nCopies(Math.max(0, r-r1), "U")) +
                     String.join("", Collections.nCopies(Math.max(0, c1-c), "R")) +
                     String.join("", Collections.nCopies(Math.max(0, c-c1), "L")) +
                     String.join("", Collections.nCopies(Math.max(0, r1-r), "D")) +
                     "!");
            r = r1;
            c = c1;
        }
        return sb.toString();
    }
}