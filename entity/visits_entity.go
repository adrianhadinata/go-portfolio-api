package entity

type Visitor struct {
	Visit_ID     int    `json:"visit_id"`
	Visitor_IP   string `json:"visitor_ip"`
	Date_Created string `json:"date_created"`
}