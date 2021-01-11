class Solution {
    public int strongPasswordChecker(String password) {
        int missingType = 3;
        boolean hasLower = false, hasUpper = false, hasDigit = false;
        for (char ch: password.toCharArray()) {
            if ('a' <= ch && ch <= 'z') {
                hasLower = true;
            } else if ('A' <= ch && ch <= 'Z') {
                hasUpper = true;
            } else if ('0' <= ch && ch <= '9') {
                hasDigit = true;
            }
        }
        if (hasLower) {
            missingType -= 1;
        }
        if (hasUpper) {
            missingType -= 1;
        }
        if (hasDigit) {
            missingType -= 1;
        }
        int replace = 0;
        int deleteOne = 0, deleteTwo = 0;
        int i = 2, n = password.length();
        while (i < n) {
            char ch = password.charAt(i);
            if (ch == password.charAt(i-1) && ch == password.charAt(i-2)) {
                int repeat = 2;
                while (i < n && ch == password.charAt(i)) {
                    i++;
                    repeat++;
                }
                replace += repeat / 3;
                if (repeat % 3 == 0) {
                    deleteOne += 1;
                } else if (repeat % 3 == 1) {
                    deleteTwo += 1;
                }
            } else {
                i++;
            }
        }
        if (n < 6) {
            return Math.max(missingType, 6-n);
        } else if (n <= 20) {
            return Math.max(missingType, replace);
        } else {
            int delete = n - 20;
            replace -= Math.min(delete, deleteOne);
            replace -= Math.min(Math.max(0, delete - deleteOne), deleteTwo * 2) / 2;
            replace -= Math.max(0, delete - deleteOne - 2 * deleteTwo) / 3;
            return delete + Math.max(missingType, replace);
        }
    }
}