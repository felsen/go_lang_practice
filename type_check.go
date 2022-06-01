
package main


import (
    "fmt"
    "time"
    "reflect"
    )

// reflect is the package which helps to identify the type of the variable.

func main() {

    firstName, lastName, age := "Felix", "Stephen", 30

    start := time.Now()
    fmt.Printf("%T\n", start)  // time.Time


    // TypeOf() is the function which helps to find the data type of the varibale.

    fmt.Println(reflect.TypeOf(start))

    fmt.Println(reflect.TypeOf(firstName))
    fmt.Println(reflect.TypeOf(lastName))
    fmt.Println(reflect.TypeOf(age))

    // ValueOf() and Kind() are the function which helps to identify the data structure of the variable.

    fmt.Println(reflect.ValueOf(firstName).Kind())
    fmt.Println(reflect.ValueOf(lastName).Kind())
    fmt.Println(reflect.ValueOf(age).Kind())

}


