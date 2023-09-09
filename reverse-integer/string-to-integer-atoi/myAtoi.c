int myAtoi(char * s){
    int i;
    int sign;
    int num;

    i = 0;
    sign = 1;
    num = 0;
    while (*s <= 32 && *s >= 13)
        s++;
    if (*s == '-')
    {
        s++;
        sign *= -1;
    }
    else if (*s == '+')
    {
        s++;
    }
    while (*s >= '0' && *s <= '9'){
        int digit = *s++ - '0';
        if (num > INT_MAX / 10 || (num == INT_MAX / 10 && digit > 7)) {
            return (sign == 1) ? INT_MAX : INT_MIN;
        }
        num = ((10 * num) + digit);
    }
    return (num * sign);
}
