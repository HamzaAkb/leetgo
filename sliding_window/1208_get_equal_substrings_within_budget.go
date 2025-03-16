import (
    "math"
)

func absoluteDiff(a, b byte) int {
    return int(math.Abs(float64(a) - float64(b)))
}

func equalSubstring(s string, t string, maxCost int) int {
    left, maxLength, cost := 0, 0, 0

    for right := 0; right < len(s); right++ {
        cost += absoluteDiff(s[right], t[right])

        for cost > maxCost {
            cost -= absoluteDiff(s[left], t[left])
            left++
        }

        if cost <= maxCost && right - left + 1 > maxLength {
            maxLength = right - left + 1
        }
    }

    return maxLength
}
