package message

type CreateEmployeeCommand struct {
	Name     string
	Email    string
	Phone    string
	Address  string
	Level    string
	Position string
}

type GetActiveEmployeeByIDQuery struct {
	ID int64
}
