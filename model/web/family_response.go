package web

type FamilyResponse struct {
	Relation string `json:"relasi"`
	Name     string `json:"nama"`
	Dob      string `json:"tanggal_lahir"`
}
