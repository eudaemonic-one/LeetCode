func replaceElements(arr []int) []int {
    if len(arr) == 0 {
        return arr
    }
    INT_MIN := -int(^uint(0)>>1)
    for i := 0; i < len(arr)-1; i++ {
        greatest := INT_MIN
        for j := i+1; j < len(arr); j++ {
            if arr[j] > greatest {
                greatest = arr[j]
                arr[i] = greatest
            }
        }
    }
    arr[len(arr)-1] = -1
    return arr
}