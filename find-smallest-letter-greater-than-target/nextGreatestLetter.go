package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
    n := len(letters)
    left, right := 0, n-1


    if target >= letters[right] {
        return letters[0]
    }


    for left <= right{
        mid := left + (right-left)/2
        if letters[mid] <= target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    if left >= n {
        left = 0
    }

    return letters[left]
}
