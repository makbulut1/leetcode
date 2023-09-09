import "math"

func myAtoi(s string) int {
    num := 0
    sign := 1
    i := 0

    for i < len(s) && s[i] == ' ' {
        i++
    }

    if i < len(s) && (s[i] == '-' || s[i] == '+') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')
        if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && digit > 7) {
            if sign == 1 {
                return math.MaxInt32
            } else {
                return math.MinInt32
            }
        }

        num = num*10 + digit
        i++
    }
    return num * sign
}
