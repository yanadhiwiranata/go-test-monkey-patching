package domain

type User struct {
	ID   int
	Name string
}

type UserFunction interface {
	GetId() int
	GetName() string
	GetIDAndName() string
}
