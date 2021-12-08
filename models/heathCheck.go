package models

import "time"

type HeathCheckStatus struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}
