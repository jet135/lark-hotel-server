package bo

import "time"

type Bill struct {
	Name                 string    `json:"name"`
	RoomNumber           string    `json:"roomNumber"`
	Phone                string    `json:"phone"`
	RoomPrice            float64   `json:"roomPrice"`
	RoomPaymentType      string    `json:"roomPaymentType"`
	Deposit              float64   `json:"deposit"`
	DepositType          string    `json:"depositType"`
	Amount               float64   `json:"amount"`
	PayTime              time.Time `json:"payTime"`
	Shift                string    `json:"shift"`
	Remark               string    `json:"remark"`
	NumberOfNights       int       `json:"numberOfNights"`
	ExpectedCheckoutTime time.Time `json:"expectedCheckoutTime"`
	CheckinTime          time.Time `json:"checkinTime"`
	CheckoutTime         time.Time `json:"checkoutTime"`
	DepositRefundTime    time.Time `json:"depositRefundTime"`
}
