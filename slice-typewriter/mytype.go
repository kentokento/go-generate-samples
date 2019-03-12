package main

import "fmt"

// +gen slice:"Average"
type MyType float64

func main() {
	temps := MyTypeSlice{15.1, -2, 3.6}

	fmt.Println(temps.Average()) // => 5.567, nil
}
