# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
    for _, letter := range letters {
        if letter > target {
            return letter
        }
    }
    return letters[0]
}
```
