package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Profile               Profile            `json:"profile" bson:"profile"`
	Emails                []Email            `json:"emails" bson:"emails"`
	ActiveOrganisationID  string             `json:"activeOrganisationId" bson:"activeOrganisationId"`
	LinkedOrganisationIDs []string           `json:"linkedOrganisationIds" bson:"linkedOrganisationIds"`
}

type Profile struct {
	FirstName      string    `json:"FirstName" bson:"FirstName"`
	LastName       string    `json:"LastName" bson:"LastName"`
	PhoneNumber    string    `json:"PhoneNumber" bson:"PhoneNumber"`
	DateOfBirth    time.Time `json:"DateOfBirth" bson:"DateOfBirth"`
	NationalNumber string    `json:"NationalNumber" bson:"NationalNumber"`
	CountryCode    string    `json:"CountryCode" bson:"CountryCode"`
	Settings       Settings  `json:"Settings" bson:"Settings"`
	Gender         string    `json:"Gender" bson:"Gender"`
	Formations     []string  `json:"Formations" bson:"Formations"`
	Hobbies        []string  `json:"Hobbies" bson:"Hobbies"`
	Experiences    []string  `json:"Experiences" bson:"Experiences"`
}

type Settings struct {
	Locale       string `json:"Locale" bson:"Locale"`
	ActiveCityID string `json:"ActiveCityID" bson:"ActiveCityID"`
}

type Email struct {
	Address  string `json:"Address" bson:"Address"`
	Verified bool   `json:"Verified" bson:"Verified"`
}
