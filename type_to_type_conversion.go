
package main


import (
    "fmt"
    "reflect"
    "strconv"
    )


func main() {

    var age int
    fmt.Print("Enter your age: ")

    // This function scans the entire input from the console, and it takes only integer. tyr out Ex: forty-five, 40-five

    fmt.Scanf("%d", &age)  // this function reads the input from the console.
    fmt.Println("You entered: ", age)
    fmt.Println(reflect.TypeOf(age))

    // strinput()

    // strbool()
 
    // strfloat()

    // strint()

    strunit()

}


func strunit() {
    u1, err := strconv.ParseUint("18", 10, 23)
    fmt.Println(u1)
    fmt.Println(err)
    fmt.Printf("%T\n", u1)

    u2, err := strconv.ParseUint("-18", 34, 5)
    fmt.Println(u2)
    fmt.Println(err)
    fmt.Printf("%T\n", u2)
}


func strint() {
    i, err := strconv.ParseInt("-18.44", 10, 64)
    fmt.Println(i)
    fmt.Println(err)
    fmt.Printf("%T\n", i)
}


func strfloat() {
    f, err := strconv.ParseFloat("3.14", 64)
    fmt.Println(f)
    fmt.Println(err)
    fmt.Printf("%T\n", f)
}


func strbool() {
    b, err := strconv.ParseBool("t")
    fmt.Println(b)
    fmt.Println(err)
    fmt.Printf("%T\n", b)
}


func strinput() {
    var input string
    fmt.Print("Enter your age: ")
    fmt.Scanf("%s", &input)

    // fmt.Println(reflect.TypeOf(input)) // this returns string data type


    // strconv is the package and Atoi is the function which helps convert string to integer
    age, err := strconv.Atoi(input)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Your age is: ", age)
    }

}

