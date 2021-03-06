package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var dat string
    var sum int64 = 0

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    check(scanner.Err())
    dat = scanner.Text()
    chars := strings.Split(dat, "")
    charlen := len(chars)

    for i := 0 ; i < charlen ; i++ {
        // the % here takes care of wrapping
        if chars[i] == chars[(i+1)%charlen] {
            x, err := strconv.ParseInt(chars[i], 10, 32)
            check(err)
            sum += x
        }
    }
    fmt.Printf("%d\n", sum)
}
