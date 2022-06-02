

package main


import (
    "fmt"
    "strconv"
    )


func main() {
    queue := 5
    name := "Felix"
    // s := name + ", your queue number is, " + queue  // This returns with error because integer cannot be converted with string.
    s := name + ", your queue number is, " + strconv.Itoa(queue)
    fmt.Println(s)

    fmt.Sprintf("%s your queue number is %d", name, queue)
}

