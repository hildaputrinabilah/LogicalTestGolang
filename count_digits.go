package main

import (
    "fmt"
    "unicode"
)

func countDigits(strArray []string) int {
    count := 0
    for _, str := range strArray {
        for _, char := range str {
            if unicode.IsDigit(char) {
                count++
            }
        }
    }
    return count
}

func main() {
    testCases := [][]string{
        {"b", "7", "h", "6", "h", "i", "5", "g", "7", "8"},
        {"7", "b", "8", "5", "5", "6", "9", "n", "f", "y", "6", "9"},
        {"u", "h", "h", "7", "6", "5", "1", "g", "7", "9"},
    }

    for _, strArray := range testCases {
        fmt.Println(countDigits(strArray))
    }
}
