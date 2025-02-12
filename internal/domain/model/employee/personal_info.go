package employee

type PersonalInfo struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

func NewPersonalInfo(name, email, phone, address string) PersonalInfo {
	return PersonalInfo{
		Name:    name,
		Email:   email,
		Phone:   phone,
		Address: address,
	}
}
