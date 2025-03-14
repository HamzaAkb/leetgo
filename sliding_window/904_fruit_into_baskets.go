func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}

func totalFruit(fruits []int) int {
    seen := make(map[int]int)
    left, maxFruits := 0, 0
    
    for right:=0; right<len(fruits); right++ {
        seen[fruits[right]]++

        for len(seen) > 2 && left < right {
            seen[fruits[left]]--
            if seen[fruits[left]] == 0 {
                delete(seen, fruits[left])
            }
            left++
        }  
        
        maxFruits = max(maxFruits, right - left + 1)
    }

    return maxFruits
}