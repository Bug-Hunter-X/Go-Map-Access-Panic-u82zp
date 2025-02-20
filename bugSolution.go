```go
package main

import (

    "fmt"
)

func main() {
    var m map[string]int
    m = make(map[string]int)
    value, ok := m["a"]
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }

    //Alternative: 
    if val, ok := m["a"]; ok{
        fmt.Println("Value:", val)
    } else {
        fmt.Println("Key not found")
    }
}```