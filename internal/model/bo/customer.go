package bo

import "time"

type Customer struct {
	Name            string    `json:"name"`
	BirthDate       string    `json:"birthDate"`
	Czdz            string    `json:"czdz"`
	IdCode          string    `json:"idCode"`
	Phone           string    `json:"phone"`
	LastCheckInTime time.Time `json:"lastCheckInTime"`
	Bills           []Bill
	BillTotal       int
}
