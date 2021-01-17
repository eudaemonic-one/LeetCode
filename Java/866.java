class Solution {
    public int primePalindrome(int N) {
        while (N != reverse(N) || !isPrime(N)) {
            N++;
            if ((int)1e7 < N && N < (int)1e8) {
                N = (int)1e8;
            }
        }
        return N;
    }
    
    private int reverse(int num) {
        int reversed = 0;
        while (num > 0) {
            reversed = 10 * reversed + (num % 10);
            num /= 10;
        }
        return reversed;
    }
    
    private boolean isPrime(int num) {
        if (num < 2) {
            return false;
        }
        int sq = (int) Math.sqrt(num);
        for (int i = 2; i <= sq; i++) {
            if (num % i == 0) {
                return false;
            }
        }
        return true;
    }
}