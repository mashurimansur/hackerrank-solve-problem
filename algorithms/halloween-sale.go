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
 * Complete the 'howManyGames' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER d
 *  3. INTEGER m
 *  4. INTEGER s
 */

func howManyGames(p int, d int, m int, s int) int {
    var total int
    var count int

    for {
        if total+p > s {
            break
        }
        total += p
        count++

        if p-d > m {
            p -= d
        } else {
            p = m
        }
    }

    return count
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16*1024*1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    pTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    p := int(pTemp)

    dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    d := int(dTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    m := int(mTemp)

    sTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
    checkError(err)
    s := int(sTemp)

    answer := howManyGames(p, d, m, s)

    fmt.Fprintf(writer, "%d\n", answer)

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
