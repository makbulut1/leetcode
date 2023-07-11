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
func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if index, ok := numMap[complement]; ok {
            return []int{index, i}
        }
        numMap[num] = i
    }

    return nil
}
```
