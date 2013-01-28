package main

import (
    "fmt"
    "math/cmplx"
)

func Cbrt(x complex128) complex128 {
    z := 1 + 0i
    for _ = range make([]int, 10) {
        z -= (z*z*z - x) / (3 * z * z)
    }
    return z
}

func main() {
    for _, n := range []complex128{1i, 2, 3 + 4i} {
        cbrtN := Cbrt(n)
        fmt.Println(cbrtN, cmplx.Pow(cbrtN, 3))
    }
}
