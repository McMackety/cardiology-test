package model

type Patient struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Patient) IsNode()               {}
func (patient Patient) GetID() string { return patient.ID }
