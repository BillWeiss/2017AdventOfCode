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

    // the last int is a special case from wrapping
    for i := 0 ; i < (len(chars)-1) ; i++ {
        if chars[i] == chars[i+1] {
            x, err := strconv.ParseInt(chars[i], 10, 32)
            check(err)
            sum += x
        }
    }

    // special wrap-around case
    if chars[len(chars)-1] == chars[0] {
        x, err := strconv.ParseInt(chars[0], 10, 32)
        check(err)
        sum += x
    }

    fmt.Printf("%d\n", sum)
}
