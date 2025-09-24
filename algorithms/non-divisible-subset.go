package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func nonDivisibleSubset(k int, s []int) int {
    // Hitung jumlah tiap remainder
    remainderCount := make([]int, k)
    for _, v := range s {
        remainderCount[v%k]++
    }

    // Minimal ambil 1 dari remainder 0 (kalau ada)
    result := 0
    if remainderCount[0] > 0 {
        result++
    }

    // Iterasi dari 1 sampai k/2
    for r := 1; r <= k/2; r++ {
        if r == k-r {
            // Kasus khusus kalau k genap (contoh: k=4, r=2)
            if remainderCount[r] > 0 {
                result++
            }
        } else {
            // Pilih kelompok dengan count lebih besar
            if remainderCount[r] > remainderCount[k-r] {
                result += remainderCount[r]
            } else {
                result += remainderCount[k-r]
            }
        }
    }

    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int(kTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int

    for i := 0; i < int(n); i++ {
        sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        sItem := int(sItemTemp)
        s = append(s, sItem)
    }

    result := nonDivisibleSubset(k, s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
