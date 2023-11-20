package o

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
}

func NewNotExtendableUser(name string) *User {
	return &User{Name: name}
}

func (u *User) PrintName() {
	fmt.Println("His name is", u.Name)

	if reflect.TypeOf(User{}).NumField() != 2 {
		panic("Only 2 fields supported!")
	}
}
