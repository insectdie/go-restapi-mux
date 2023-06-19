package domain

type Customer struct {
	Id          int
	Name        string
	Dob         string
	Phone       string
	Nationality string
	Email       string
	Family      interface{}
}
