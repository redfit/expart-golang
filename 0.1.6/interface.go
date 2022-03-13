package main

import "fmt"

type Crier interface {
	Cry() string
}

type Duck struct{}

func (d *Duck) Cry() string {
	return "Quack"
}

type ParrotFunc func() string

func (p ParrotFunc) Cry() string {
	return p()
}

func main() {

	var d Crier = &Duck{}

	fmt.Println(d.Cry())

	var p Crier = ParrotFunc(func() string {
		return "Squawk"
	})

	fmt.Println(p.Cry())

	var _ Crier = (*ParrotFunc)(nil)

	ss := []string{"John", "Richard"}

	//var i interface{} = ss // OK
	//var is []interface{} = ss // NG
	is := make([]interface{}, len(ss))
	for _, s := range ss {
		is = append(is, s)
	}

	fmt.Println(is)

}
