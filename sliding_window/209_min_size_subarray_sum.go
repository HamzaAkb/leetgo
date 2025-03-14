func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func minSubArrayLen(target int, nums []int) int {
    left, minLength, currentSum := 0, len(nums) + 1, 0

    for right, val := range nums {
        currentSum += val

        for currentSum >= target {
            minLength = min(minLength, right-left+1)
            currentSum -= nums[left]
            left++
        }
    }

    if minLength == len(nums) + 1 { return 0 }
    
    return minLength
}
