package models

import "time"

type Todo struct {
	ID        uint      `json:”id”`
	name      string    `json:”name”`
	startDate time.Time `json:”startdate”`
	endDate   time.Time `json:”enddate”`
}
