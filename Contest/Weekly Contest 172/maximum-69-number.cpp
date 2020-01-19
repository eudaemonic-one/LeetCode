class Solution {
public:
    int maximum69Number (int num) {
        int maximum = num;
        int digits = 0;
        int tmp = num;
        while (tmp > 0) {
            tmp /= 10;
            digits++;
        }
        for (int i = digits-1; i >= 0; i--) {
            int tmp = num % int(pow(10, i+1));
            if (tmp / int(pow(10, i)) == 6) {
                maximum += 3 * pow(10, i);
                break;
            }
        }
        return maximum;
    }
};