func decompressRLElist(nums []int) []int {
    decompressed := make([]int, 0)
    for i := 0; i < len(nums)-1 && 2*i+1 < len(nums); i++ {
        a := nums[2*i]
        b := nums[2*i+1]
        for j := 0; j < a; j++ {
            decompressed = append(decompressed, b)
        }
    }
    return decompressed
}