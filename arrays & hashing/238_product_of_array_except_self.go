func productExceptSelf(nums []int) []int {
    result := make([]int, len(nums))
    productRight := make([]int, len(nums))
    arrayLength := len(nums)

    result[0] = 1
    for i:=1; i<arrayLength; i++ {
        result[i] = result[i - 1] * nums[i - 1]
    }

    productRight[arrayLength - 1] = 1
    for i:=arrayLength - 2; i>=0; i-- {
        productRight[i] = productRight[i + 1] * nums[i + 1]
        result[i] = productRight[i] * result[i]
    }

    return result
}