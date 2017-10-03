package memberModel

type Family struct {
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip int `json:"zip"`
	Adults []Member `json:"adults"`
	Children []Member `json:"children"`
}

type Member struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	CellPhone string `json:"cell_phone"`
	BirthDate string `json:"birthDate"`
}

type Families struct {
	Families []Family
}

