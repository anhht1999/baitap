package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {

    words := []string{"sky", "florest", "fly", "cup", "wood",
        "falcon", "so", "see", "tool"}
	//tìm theo chữ
    filtered := Filter(words, func(word string) bool {
        return strings.HasPrefix(word, "fl")
    })

    fmt.Println(filtered)
	//tim theo do dai chuỗi
    filtered2 := Filter(words, func(word string) bool {
        return len(word) == 3
    })

    fmt.Println(filtered2)

    
    vals := []int{-2, 0, 1, 9, 7, -3, -5, 6}
    positive := []int{}

    for i := range vals {
        if vals[i] > 0  {
            positive = append(positive, vals[i])
        }
    }

    fmt.Println(positive)

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter a number: ")
    str1, _ := reader.ReadString('\n')

    // remove newline
    str1 = strings.Replace(str1,"\n","",-1)

    // convert string variable to int variable
    num, e := strconv.Atoi(str1)
    if e != nil {
	fmt.Println("conversion error:", str1)
    }
    if num >= 1 && num <= 10 {
        fmt.Println("correct")
        } else {
            fmt.Println("num not in range")
        }
}

func Filter(vs []string, f func(string) bool) []string {
    filtered := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            filtered = append(filtered, v)
        }
    }
    return filtered
}