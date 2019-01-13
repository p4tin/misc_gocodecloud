package main

import "fmt"

func SwapRune(r rune) rune {
    switch {
    case 97 <= r && r <= 122:
        return r - 32
    case 65 <= r && r <= 90:
        return r + 32
    default:
        return r
    }
}

func main() {
    fmt.Println(SwapRune('a'))
}

