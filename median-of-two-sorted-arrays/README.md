# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
The code aims to find the median of two sorted arrays, nums1 and nums2. It combines the two arrays into a single sorted array and calculates the median based on the length of the combined array.
# Approach
<!-- Describe your approach to solving the problem. -->

1. Initialize variables i, j, and k to keep track of the indices of nums1, nums2, and resultArray, respectively.
2. Determine the lengths of nums1 and nums2 and store them in variables m and n.
3. Create an array called resultArray with a size of m + n to store the merged and sorted values from nums1 and nums2.
4. Use a while loop to merge the arrays until either nums1 or nums2 is fully processed.
5. Compare the elements at indices i and j of nums1 and nums2. Add the smaller element to resultArray and increment the corresponding index (i or j) and k.
6. Once the merging is complete, calculate the median of resultArray based on its length (m + n).
7. If the combined length is even, take the average of the middle two elements as the median.
8. If the combined length is odd, the median is the middle element.
9. Return the calculated median.

# Complexity
- Time complexity:The code has a time complexity of O(m + n) since it iterates through both nums1 and nums2 exactly once to merge the arrays.
Space complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: The code has a space complexity of O(m + n) because it creates resultArray with a size of m + n to store the merged elements.
Overall, the code takes two sorted arrays, merges them into a single sorted array, and calculates the median based on the combined array's length.
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    i, j, k := 0, 0, 0;

    m := len(nums1);
    n := len(nums2);

    resultArray := make([]int, m + n)
 
    for i < len(nums1) || j < len(nums2) {
        if j >= len(nums2) {
            resultArray[k] = nums1[i]
            i++
        } else if i >= len(nums1) {
            resultArray[k] = nums2[j]
            j++
        } else if nums1[i] < nums2[j] {
            resultArray[k] = nums1[i]
            i++
        } else {
            resultArray[k] = nums2[j]
            j++
        }
        k++
    }
    var ret float64
    if (n + m) % 2 == 0 {
        data := ((n + m) / 2) - 1
        ret1 := resultArray[data]
        ret2 := resultArray[data + 1]
        ret = (float64(ret1) + float64(ret2)) / 2.0
    } else {
        data := ((n + m) / 2)
        ret = float64(resultArray[data])
    }
    return ret;
}
```
