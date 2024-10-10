package types

type UserType int

const (
	UserTypeUnknown   = 0
	UserTypeVendedor  = 1
	UserTypeComprador = 2
)

type User struct {
	Id       uint
	Name     string
	Cpf      string
	Email    string
	Password string
	UsrType  UserType
}
