## Intuition
Explain your initial thoughts on how to solve this problem.

## Approach
Explain your approach to solving the problem.

## Complexity
- Time complexity:
  - Add your time complexity here, e.g., O(n)

- Space complexity:
  - Add your space complexity here, e.g., O(n)

## Code
```go
func length(x int) int {
  len := 0
  for ; x > 0; len++ {
    x = x / 10
  }
  return len
}

func itoa(x int) string {
    var len int
    var ret []byte;
    tmp := x
    if x < 0{
        x = x * -1
        len = length(x) + 1
        ret = make([]byte, len)    
    } else{
        len = length(x)
        ret = make([]byte, len)
    }
    for i := len - 1; i >= 0; i-- {
        ret[i] = byte((x % 10) + 48)
        x = x / 10
    }
    if (tmp < 0){
        ret[0] = '-'
    }
    return string(ret)
}

func rev(str string) string {
    ret := make([]byte, len(str))
    i := 0
    for len := len(str); len > 0; len-- {
        ret[i] = str[len - 1]
        i++
    }
    return string(ret)
}

func isPalindrome(x int) bool {
    array := itoa(x)
    rearray := rev(itoa(x))
    if array == rearray {
        return true
    }
    return false
}
