func cuttingRope(n int) int {
    if n < 4 {
        return n-1
    }
    res := 1 
    for n>4{
        res = res * 3 % 1000000007
        n-=3
    }
    return res * n % 1000000007
}

// 取余数
// 贪心数学推导