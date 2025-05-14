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
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func beautifulDays(i int, j int, k int) int {
    // Write your code here
    data := []int{}
    for in := i; in <= j; in++ {
        // fmt.Println(in)
        // data = append(data, in)
        // fmt.Println(in)
        reverse := 0
        num := in
        for num > 0 {
            reverse = reverse*10 + num%10
            num /= 10
        }
        if (in-reverse)%k == 0 {
            data = append(data, in)
            continue
        }
    }

    return len(data)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    i := int(iTemp)

    jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    j := int(jTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    k := int(kTemp)

    result := beautifulDays(i, j, k)

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
