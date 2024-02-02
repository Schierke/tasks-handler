package models

import "time"

type TaskResult struct {
	ID           string `json:"id" bson:"_id"`
	Name         string `json:"name" bson:"name"`
	Organisation struct {
		Name       string `json:"name" bson:"name"`
		Address    string `json:"address" bson:"address"`
		PictureUrl string `json:"pictureurl" bson:"pictureurl"`
	} `json:"organisation" bson:"organisation"`
	Slots struct {
		Filled int `json:"filled" bson:"filled"`
		Total  int `total:"filled" bson:"total"`
	} `json:"slots" bson:"slotsdata"`
}

type TaskResult2 struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Users struct {
		FirstName string `json:"first_name" bson:"firstname"`
		LastName  string `total:"first_name" bson:"lastname"`
	} `json:"users" bson:"userdata"`
}

type Task struct {
	ID                   string          `bson:"_id" json:"_id"`
	LogAsID              string          `bson:"logAsId" json:"logAsId"`
	CreatorID            string          `bson:"creatorId" json:"creatorId"`
	OrganisationID       string          `bson:"organisationId" json:"organisationId"`
	ShiftIDs             []string        `bson:"shiftIds" json:"shiftIds"`
	CityID               string          `bson:"cityId" json:"cityId"`
	ManagerID            string          `bson:"managerId" json:"managerId"`
	RequestedSiderIDs    []string        `bson:"requestedSiderIds" json:"requestedSiderIds"`
	AssigneeID           string          `bson:"assigneeId" json:"assigneeId"`
	Alias                string          `bson:"alias" json:"alias"`
	Status               string          `bson:"status" json:"status"`
	Address              Address         `bson:"address" json:"address"`
	LocationOptions      LocationOptions `bson:"locationOptions" json:"locationOptions"`
	WorkLegalStatus      string          `bson:"workLegalStatus" json:"workLegalStatus"`
	SelectionStatus      string          `bson:"selectionStatus" json:"selectionStatus"`
	RequestedSidersOnly  bool            `bson:"requestedSidersOnly" json:"requestedSidersOnly"`
	IsPreSelection       bool            `bson:"isPreSelection" json:"isPreSelection"`
	Visible              bool            `bson:"visible" json:"visible"`
	UsersAlreadyNotified bool            `bson:"usersAlreadyNotified" json:"usersAlreadyNotified"`
	CompanyNotified      bool            `bson:"companyNotified" json:"companyNotified"`
	Type                 string          `bson:"type" json:"type"`
	SubtaskIDs           []string        `bson:"subtaskIds" json:"subtaskIds"`
	Purpose              string          `bson:"purpose" json:"purpose"`
	DressCode            string          `bson:"dressCode" json:"dressCode"`
	WorkConditions       string          `bson:"workConditions" json:"workConditions"`
	Experiences          string          `bson:"experiences" json:"experiences"`
	Motive               Motive          `bson:"motive" json:"motive"`
	MissionInformation   string          `bson:"missionInformation" json:"missionInformation"`
	SideNote             string          `bson:"sideNote" json:"sideNote"`
	PricingID            string          `bson:"pricingId" json:"pricingId"`
	HourlyRate           float64         `bson:"hourlyRate" json:"hourlyRate"`
	SubmittedAt          time.Time       `bson:"submittedAt" json:"submittedAt"`
	LiveAt               time.Time       `bson:"liveAt" json:"liveAt"`
	PostedAt             time.Time       `bson:"postedAt" json:"postedAt"`
}

// Address represents the address field in the tasks document
type Address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	ZipCode string `bson:"zipCode" json:"zipCode"`
	Country string `bson:"country" json:"country"`
}

// LocationOptions represents the locationOptions field in the tasks document
type LocationOptions struct {
	Remote    RemoteLocation `bson:"remote" json:"remote"`
	Motorized bool           `bson:"motorized" json:"motorized"`
}

// RemoteLocation represents the remote field in the locationOptions
type RemoteLocation struct {
	Available bool `bson:"available" json:"available"`
	Mandatory bool `bson:"mandatory" json:"mandatory"`
}

// Motive represents the motive field in the tasks document
type Motive struct {
	Reason       string        `bson:"reason" json:"reason"`
	Replacements []Replacement `bson:"replacements" json:"replacements"`
}

// Replacement represents a replacement in the motive field
type Replacement struct {
	Name          string `bson:"name" json:"name"`
	Position      string `bson:"position" json:"position"`
	Justification string `bson:"justification" json:"justification"`
}
