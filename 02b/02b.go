package main

import (
    "bufio"
    "fmt"
//    "math"
    "os"
    "sort"
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
        check(scanner.Err())
        dat := scanner.Text()
        dats := strings.Fields(dat)

        // turn the array into ints
        ints := make([]int, len(dats))
        for i := range dats {
            ints[i], _ = strconv.Atoi(dats[i])
        }

        // by sorting the array, we can only check elements after the
        // current one for division
        sort.Sort(sort.Reverse(sort.IntSlice(ints)))

        for i := range ints {
            for j := i ; j < len(ints) ; j++ {
                if ints[j] < ints[i] {
                    if ints[i] % ints[j] == 0 {
                        checksum += ints[i] / ints[j]
                        // I feel weird about not stopping execution here,
                        // but they promise there will only be one pair
                        // like this in each row
                    }
                }
            }
        }
    }

    fmt.Printf("Checksum: %d\n", checksum)
}
