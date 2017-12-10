package main

import (
    "bufio"
    "fmt"
    "math"
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
    var checksum int = 0

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var min, max int = math.MaxInt32, math.MinInt32
        check(scanner.Err())
        dat := scanner.Text()
        dats := strings.Fields(dat)
        for i := range dats {
            tmp, _ := strconv.Atoi(dats[i])
            if tmp < min {
                min = tmp
            }
            if tmp > max {
                max = tmp
            }
        }

        checksum += max - min
    }

    fmt.Printf("%d\n", checksum)
}
