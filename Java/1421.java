class Solution {
    public int maxDiff(int num) {
        char[] a = Integer.toString(num).toCharArray(), b = a.clone();
        char x = a[0], y = 0;
        for (int i = 0; i < a.length; i++) {
            if (a[i] == x) {
                a[i] = '9';
                b[i] = '1';
            }else if (x == '1' && a[i] > '0' || x == '9' && a[i] < '9') {
                if (y == 0) {
                    y = a[i];
                } 
                if (y == a[i]) {
                    if (x == '1')
                        b[i] = '0';
                    else
                        a[i] = '9';
                }
            }
        }
        return Integer.parseInt(String.valueOf(a)) - Integer.parseInt(String.valueOf(b));
    }
}