package main

import "fmt"

func main() {
	var i int32 = 100
	var j int64

	j = int64(i)
	fmt.Println(j)

	msg := "Go Expert"
	bs := []byte(msg)
	fmt.Println(bs)

	s := string(bs)
	fmt.Println(s)

	g := interface{}("Go Expert")
	s = g.(string)
	fmt.Println(s)

	//n := g.([]byte)
	n, ok := g.([]byte)
	fmt.Println(n, ok)

	switch g.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("This is integer", g)
	case string:
		fmt.Println("This is string", g)
	default:
		fmt.Printf("This is unknown type, %T\n", g)
	}

	type ErrNoSuchEntity struct{ error }
	type ErrConflictEntity struct{ error }

	do := func() error {
		return &ErrNoSuchEntity{}
	}

	switch do().(type) {
	case *ErrNoSuchEntity:
		fmt.Println("This is ErrNoSuchEntity")
	case *ErrConflictEntity:
		fmt.Println("This is ErrConflictEntity")
	default:
		fmt.Println("This is unknown type")
	}
}
