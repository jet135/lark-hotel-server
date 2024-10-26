package report

import "time"

type Req struct {
	OriginalUrl string      `json:"originalUrl"`
	Method      string      `json:"method"`
	RequestBody interface{} `json:"requestBody"`
}

type ServiceDescript struct {
	Data          map[string]interface{} `json:"data"`
	ServiceMethod string                 `json:"serviceMethod"`
	ServiceName   string                 `json:"serviceName"`
}

type Message struct {
	ServiceDescript ServiceDescript `json:"serviceDescript"`
	ServiceID       string          `json:"serviceID"`
}

type Data struct {
	Name       string `json:"xm"`
	BirthDate  string `json:"csrq"`
	Czdz       string `json:"czdz"`
	RoomNumber string `json:"fjh"`
	IdCode     string `json:"idCode"`
	IdType     string `json:"idType"`
	Phone      string `json:"lxdh"`
	XzqhTitle  string `json:"xzqhTitle"`
	Date       time.Time
	Shift      string
}
