class Solution {
    public int numSteps(String s) {
        int steps = 0, carry = 0;
        for (int i = s.length()-1; i > 0; i--) {
            steps++;
            if (s.charAt(i) - '0' + carry == 1) {
                ++steps;
                carry = 1;
            }
        }
        return steps + carry;
    }
}