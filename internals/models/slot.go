package models

import "time"

// Slot represents the slots document in MongoDB
type Slot struct {
	ID             string    `bson:"_id" json:"_id"`
	TaskID         string    `bson:"taskId" json:"taskId"`
	ShiftID        string    `bson:"shiftId" json:"shiftId"`
	OrganisationID string    `bson:"organisationId" json:"organisationId"`
	IsOverbooking  bool      `bson:"isOverbooking" json:"isOverbooking"`
	Status         string    `bson:"status" json:"status"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
	SiderStatus    string    `bson:"siderStatus" json:"siderStatus"`
}
