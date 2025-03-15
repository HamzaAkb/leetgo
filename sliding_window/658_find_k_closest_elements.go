import (
    "sort"
    "math"
)

func findClosestElements(arr []int, k int, x int) []int {
	j := sort.SearchInts(arr, x)
	left, right, length, i := j-1, j, len(arr), 0

	if right == 0 {
		return arr[:k]
	}

	if right == length {
		return arr[length-k:]
	}

	var closest []int

	for i < k && left >= 0 && right < len(arr) {
		leftDist := math.Abs(float64(arr[left] - x))
		rightDist := math.Abs(float64(arr[right] - x))

		if leftDist <= rightDist {
			closest = append(closest, arr[left])
			left--
		} else {
			closest = append(closest, arr[right])
			right++
		}
		i++
	}

	for i < k && left >= 0 {
		closest = append(closest, arr[left])
		left--
		i++
	}

	for i < k && right < len(arr) {
		closest = append(closest, arr[right])
		right++
		i++
	}

	sort.Ints(closest)

	return closest
}