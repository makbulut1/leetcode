# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
The problem requires finding the longest palindrome substring in a given string. A palindrome is a string that reads the same forwards and backward.

# Approach
<!-- Describe your approach to solving the problem. -->
The approach to solve this problem is to use a technique called "Expand Around Center." The idea is to iterate through each character in the string and consider it as the center of a potential palindrome. We then expand around the center to check if the substring is a palindrome.

The steps for the approach are as follows:

1. If the input string has a length less than 2, it is already a palindrome, so we return the string as the result.
2. Initialize variables start and end to store the indices of the longest palindrome substring found so far.
3. Iterate through each character in the string using a loop variable i:
    - Call the expandAroundCenter function twice to check for palindromes with s[i] as the center (one case where the center is a single character and another where the center is between two characters).
    - Compare the lengths of the palindromes found with the current longest palindrome. If the new palindrome is longer, update start and end accordingly.
4. Return the longest palindrome substring found by slicing the original string using start and end indices.

The expandAroundCenter function checks if a substring is a palindrome by expanding from a given center point. It takes the string s and the left and right indices as inputs. It continuously expands the substring towards both ends while the characters at the left and right indices are equal. It returns the length of the palindrome substring found.

The max function returns the maximum of two integers.
# Complexity
- Time complexity: The algorithm iterates through the string once with a linear loop, resulting in a time complexity of O(n), where n is the length of the input string.
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: The algorithm uses a constant amount of extra space, resulting in a space complexity of O(1).
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    start, end := 0, 0

    for i := 0; i < len(s); i++ {
        len1 := expandAroundCenter(s, i, i)
        len2 := expandAroundCenter(s, i, i+1)
        maxLen := max(len1, len2)

        if maxLen > end-start {
            start = i - (maxLen-1)/2
            end = i + maxLen/2
        }
    }
    return s[start : end + 1]
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```
