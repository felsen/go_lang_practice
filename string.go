package main


import "fmt"


func main() {
    var firstName string = "felix Stephen"
    fmt.Println(firstName)

    address := "The white house\n 1600 Avenew\nNew Delhi."
    fmt.Println(address)

    quotation := `"Anyone who has never made a mistake
    has never tried anything new.
    -- Felix Stephen"`
    fmt.Println(quotation)
    fmt.Println(len(quotation))

}

