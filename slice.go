package main

import (
	"fmt"
)

func Slice(a []string, nbrs ...int) []string{
    var result []string
    
    if len(nbrs) == 0 {
        return a
    }

    num1 := nbrs[0]
   
    if len(nbrs) == 1 {
        if num1 >= 0 {
            result = a[num1:]
            return result
        } else {
            num1 = len(a) + num1
            if num1 < 0 {
                num1 = 0
            } 
            result = a[num1:]
            return result
        }
    }

    if len(nbrs) == 2 {
        num2 := nbrs[1]

        if num1 > 0 && num2 >= 0 {
            if num1 > num2 {
                return nil
            }
            if num2 > len(a) {
                num2 = len(a)
            }
            result = a[num1 : num2]
            return result
        } else {
            num1 = len(a) + num1
            num2 = len(a) + num2

            result = a[num1:num2]
            return result
        }
    }

    return result
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}

    fmt.Printf("%#v\n", Slice(a, 8))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}

/*
slice
Instructions
The function receives a slice of strings and one or more integers, and returns a slice of strings. 
The returned slice is part of the received one but cut from the position indicated in the first int, 
until the position indicated by the second int.

In case there only exists one int, the resulting slice begins in the position indicated by the int and ends at 
the end of the received slice.

The integers can be negative.

Expected function
func Slice(a []string, nbrs... int) []string{

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"

	"piscine"
)

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", piscine.Slice(a, 1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
    fmt.Printf("%#v\n", piscine.Slice(a, -3))
    fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
}
$ go run .
[]string{"algorithm", "ascii", "package", "golang"}
[]string{"ascii", "package"}
[]string{"ascii", "package", "golang"}
[]string{"package"}
[]string(nil)
*/