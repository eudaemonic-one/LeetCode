/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    l, r := 0, n
    for l < r {
        m := l + (r-l)/2
        if isBadVersion(m) {
            r = m
        } else {
            l = m+1
        }
    }
    return l
}
