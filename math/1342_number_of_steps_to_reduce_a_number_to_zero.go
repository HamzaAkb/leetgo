func numberOfSteps(num int) int {
    if num <= 1 {
        if num == 0 {
            return 0
        }
        return 1
    }

    if num % 2 == 0 {
        return 1 + numberOfSteps(num / 2)
    }

    return 1 + numberOfSteps(num - 1)
}