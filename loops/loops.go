package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    for Δ := 0.1; math.Abs(Δ) > 1e-10; Δ = (z*z - x) / (2 * z) {
        z -= Δ
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
