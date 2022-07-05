package main

import "fmt"

type User interface {
	GetRole() string
}
type Member struct{}

func (t *Member) GetRole() string {
	return "普通用户"
}

type Admin struct{}

func (t *Admin) GetRole() string {
	return "管理用户"
}

const (
	Mem = iota
	Adm
)

func CreateUser(t int) User {
	switch t {
	case Mem:
		return new(Member)
	case Adm:
		return new(Admin)
	default:
		return new(Member)
	}
}

func main() {
	fmt.Println(CreateUser(Adm).GetRole())
}
