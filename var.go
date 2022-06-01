package main


import "fmt"


var num = 20

func main() {
    var num1 = 10  // in case any variable is unused that will throw an error.

    var num2 int  // integer data type declaration
    var num3 float32  // float data type declaration
    var num4 bool  // boolean data type declaration
    var name string  // string data type declaration

    var rates float32 = 4.2  // float data type declation

    fmt.Println(num)
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(num3)
    fmt.Println(num4)
    fmt.Println(name)
    fmt.Println(rates)

    short_variable_declaration()

}

func short_variable_declaration() {
    // short variable declaration

    firstName := "Felix"
    fmt.Println(firstName)

    firstName, lastName, age := "Felix", "Stephen", "30"
    fmt.Println(firstName, lastName, age)

    var fname, lname string = "felix", "stephen"
    fmt.Println(fname, lname)


    // All of these are different combination of variables, then below is not right to use, this will always gives an error.

    // var fm string, lm string, ag int = "felix", "stephen", 40
    // fmt.Println(fm, lm, ag)

    var fN, lN, aG = "felix", "stephen", 30
    fmt.Println(fN, lN, aG)

    var (
        first_name = "felix"
        last_name = "stephen"
        ags = 40
    )
    fmt.Println(first_name, last_name, ags)

}



