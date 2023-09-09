func reverse(x int) int {
    str := strconv.Itoa(x)
    newStr := make([]byte, len(str))
    neg := 0
    for i,j := (len(str) - 1), 0; (i >= 0 && j < len(str)) || (neg == 1 && i > 0); i-- {
            if str[j] == '-' {
                newStr[0] = str[j]
                neg = 1
                j++
            }
        newStr[i] = str[j];
        j++
    }
    x, err := strconv.Atoi(string(newStr))
    if err != nil || (x > 2147483647 || x < -2147483648){
        return 0
    }
    return x
}
