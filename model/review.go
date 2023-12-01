package model

import "time"

type Review struct {
	ID       int64     `json:"id"`
	Rate     int64     `json:"rate"`
	Datetime time.Time `json:"datetime"`
	User     User      `json:"user"`
	Product  Product   `json:"product"`
}
