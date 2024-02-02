package models

import "time"

// Shift represents the shifts document in MongoDB
type Shift struct {
	ID                string    `bson:"_id" json:"_id"`
	AvailableSiderIds []string  `bson:"availableSiderIds" json:"availableSiderIds"`
	HiredSiderIds     []string  `bson:"hiredSiderIds" json:"hiredSiderIds"`
	Time              ShiftTime `bson:"time" json:"time"`
	Type              string    `bson:"type" json:"type"`
	Slots             int       `bson:"slots" json:"slots"`
	TaskID            string    `bson:"taskId" json:"taskId"`
	Break             int       `bson:"break" json:"break"`
	Location          string    `bson:"location" json:"location"`
	CreatedAt         time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time `bson:"updatedAt" json:"updatedAt"`
	Status            string    `bson:"status" json:"status"`
}

// ShiftTime represents the time field in the shifts document
type ShiftTime struct {
	StartDate time.Time `bson:"startDate" json:"startDate"`
	EndDate   time.Time `bson:"endDate" json:"endDate"`
}

type ShiftResponse struct {
	ID   string    `json:"id"`
	Time ShiftTime `json:"time"`
	Filled int `json:"filled"`
	Available int `json:"available"`
}

