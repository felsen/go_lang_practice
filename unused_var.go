package main

import (
    "fmt"
    )


// if unused varibale is present in the file, it might throw an error. in case if require to present mention below like this.


func main() {
    var num = 5
    _ = num
    fmt.Println("Hello unused variable")
}

