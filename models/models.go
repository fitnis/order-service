package models

type OrderRequest struct {
	PatientID string `json:"patientId"`
	Item      string `json:"item"`
	Priority  string `json:"priority"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
