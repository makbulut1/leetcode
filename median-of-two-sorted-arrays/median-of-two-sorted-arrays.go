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
