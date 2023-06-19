package web

type CustomerResponse struct {
	Id          int         `json:"id"`
	Name        string      `json:"nama"`
	Dob         string      `json:"tanggal_lahir"`
	Phone       string      `json:"telepon"`
	Nationality string      `json:"kewarganegaraan"`
	Email       string      `json:"email"`
	Family      interface{} `json:"keluarga"`
}
