// Approach #2: Backtracking with Memo
class Solution {
    public boolean canWin(String s) {
        if (s == null || s.length() < 2) {
            return false;
        }
        Map<String, Boolean> memo = new HashMap<>();
        return backtracking(s, memo);
    }
    
    private boolean backtracking(String s, Map<String, Boolean> memo) {
        if (memo.containsKey(s)) {
            return memo.get(s);
        }
        for (int i = 0; i < s.length()-1; i++) {
            if (s.charAt(i) == '+' && s.charAt(i+1) == '+') {
                String opponent = s.substring(0, i) + "--" + s.substring(i+2);
                if (!backtracking(opponent, memo)) {
                    memo.put(s, true);
                    return true;
                }
            }
        }
        memo.put(s, false);
        return false;
    }
}

// Approach #1: Backtracking
// class Solution {
//     public boolean canWin(String s) {
//         if (s == null || s.length() < 2) {
//             return false;
//         }
//         return backtracking(s);
//     }
    
//     private boolean backtracking(String s) {
//         for (int i = 0; i < s.length()-1; i++) {
//             if (s.charAt(i) == '+' && s.charAt(i+1) == '+') {
//                 String opponent = s.substring(0, i) + "--" + s.substring(i+2);
//                 if (!backtracking(opponent)) {
//                     return true;
//                 }
//             }
//         }
//         return false;
//     }
// }