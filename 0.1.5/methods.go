package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Aging() {
	u.Age++
}

func (u User) AgingButNothingHappen() {
	u.Age++
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	u := User{Name: "foo", Age: 10}
	u.Aging()
	fmt.Println(u.Age)

	// ポインタじゃない方のmethodで++しても影響が無い
	u.AgingButNothingHappen()
	fmt.Println(u.Age)

	fv := Hex(1024).String()
	fmt.Println(fv)

	// methodを渡して後で実行することができる
	fe := Hex(1024).String
	fmt.Println(fe())
}
