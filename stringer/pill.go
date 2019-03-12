package main

import "fmt"

//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate stringer -type=Pill -output=pill_gen.go
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func main() {
	fmt.Println(Placebo.String())
	fmt.Println(Aspirin.String())
	fmt.Println(Ibuprofen.String())
	fmt.Println(Paracetamol.String())
	fmt.Println(Acetaminophen.String())
}
