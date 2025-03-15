import (
	"math"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	j := sort.SearchInts(arr, x)
	left, right := j-1, j
	n := len(arr)

	if right == 0 {
		return arr[:k]
	}
	if right == n {
		return arr[n-k:]
	}

	closest := make([]int, 0, k)

	for len(closest) < k {
		if left >= 0 && (right >= n || math.Abs(float64(arr[left]-x)) <= math.Abs(float64(arr[right]-x))) {
			closest = append(closest, arr[left])
			left--
		} else if right < n {
			closest = append(closest, arr[right])
			right++
		}
	}

	sort.Ints(closest)
	return closest
}