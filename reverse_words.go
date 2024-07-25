package main

import (
    "fmt"
    "strings"
)

func reverseWords(sentence string) string {
    words := strings.Split(sentence, " ")
    for i, word := range words {
        letters := []rune(word)
        for j, n := 0, len(letters); j < n/2; j++ {
            letters[j], letters[n-1-j] = letters[n-1-j], letters[j]
        }
        words[i] = string(letters)
    }
    return strings.Join(words, " ")
}

func main() {
    sentences := []string{
        "italem irad irigayaj",
        "iadab itsap ullareb",
        "nalub kusutret gnalali",
    }

    for _, sentence := range sentences {
        fmt.Println(reverseWords(sentence))
    }
}
