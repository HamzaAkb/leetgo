func binarySearch(nums []int, left, right, target int) int {
    if left >= right {
        if nums[left] == target {
            return left
        }
        return -1
    }

    mid := int((left + right) / 2)
    
    if nums[mid] == target {
        return mid
    } else if target > nums[mid]{
        return binarySearch(nums, mid + 1, right, target)
    } else {
        return binarySearch(nums, left, mid - 1, target)
    }
}

func search(nums []int, target int) int {
    return binarySearch(nums, 0, len(nums) - 1, target)    
}