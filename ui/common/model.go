package common

import "time"

type Server struct {
	ID          int    `json:"id"`
	TYPE        string `json:"_type"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Net         string `json:"net"`
	PingLatency string `json:"pingLatency"`
}
type Subscription struct {
	Remarks string    `json:"remarks,omitempty"`
	ID      int       `json:"id"`
	TYPE    string    `json:"_type"`
	Host    string    `json:"host"`
	Address string    `json:"address"`
	Status  time.Time `json:"status"`
	Info    string    `json:"info"`
	Servers []Server  `json:"servers"`
}
