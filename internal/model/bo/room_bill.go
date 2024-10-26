package bo

import "time"

type RoomBill struct {
	RoomNumber         string    `json:"roomNumber"`
	TotalCheckIns      int       `json:"totalCheckIns"`
	CheckInsLast30Days int       `json:"checkInsLast30Days"`
	LastCheckInTime    time.Time `json:"lastCheckInTime"`
	LowestRoomPrice    float64
	HighestRoomPrice   float64
	Bills              []Bill
	BillTotal          int
}
