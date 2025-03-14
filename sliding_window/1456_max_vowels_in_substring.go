func isVowel(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxVowels(s string, k int) int {
    left, vowelCount, maxVowels := 0, 0, 0

    for right := 0; right < k; right++ {
        if isVowel(s[right]) {
            vowelCount++
        }
    }
    maxVowels = vowelCount

    for right := k; right < len(s); right++ {
        if isVowel(s[right]) {
            vowelCount++
        }
        
        if isVowel(s[left]) {
            vowelCount--
        }

        left++
        maxVowels = max(maxVowels, vowelCount)
    }

    return maxVowels
}