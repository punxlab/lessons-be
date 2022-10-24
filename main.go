package main

import "fmt"

type T struct {
	I int
	S string
}

type T2 struct {
	F float32
}

func main() {
	t := T{
		I: 1,
		S: "str",
	}

	t2 := T2{
		F: 1.2,
	}

	fmt.Printf("t=%v t2=%v HUY\n", t, t2)
}
